package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/skye-z/betax-blog/model"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

type CommonService struct {
	Engine *xorm.Engine
	Cache  *cache.Cache
}

func (service CommonService) GetWebManifest(ctx *gin.Context) {
	check := util.GetString("github.bind")
	if check != "" {
		info, found := service.Cache.Get("manifest")
		if !found {
			var checkUser User
			err := json.Unmarshal([]byte(check), &checkUser)
			if err != nil {
				util.ReturnMessage(ctx, false, "获取博主信息失败")
			} else {
				manifest := map[string]interface{}{
					"name":             checkUser.Nickname + " Blog",
					"short_name":       checkUser.Nickname + " Blog",
					"start_url":        "/app/",
					"display":          "standalone",
					"background_color": "#F5F7F9",
					"lang":             "en",
					"scope":            "/app/",
					"description":      checkUser.Bio,
					"theme_color":      "#f5f7f9",
					"icons": []map[string]string{
						{"src": "/app/res/pwa-192x192.png", "sizes": "192x192", "type": "image/png", "purpose": "any"},
						{"src": "/app/res/pwa-512x512.png", "sizes": "512x512", "type": "image/png", "purpose": "any"},
						{"src": "/app/res/pwa-maskable-192x192.png", "sizes": "192x192", "type": "image/png", "purpose": "maskable"},
						{"src": "/app/res/pwa-maskable-512x512.png", "sizes": "512x512", "type": "image/png", "purpose": "maskable"},
					},
				}
				if checkUser.Bio == "" {
					manifest["description"] = "This is " + checkUser.Nickname + "`s Blog"
				}
				ctx.Header("Content-Type", "application/manifest+json")
				service.Cache.Set("manifest", manifest, cache.NoExpiration)
				ctx.JSON(200, manifest)
			}
		} else {
			ctx.Header("Content-Type", "application/manifest+json")
			ctx.JSON(200, info)
		}
	} else {
		util.ReturnMessage(ctx, false, "系统尚未初始化")
	}
}

type InitData struct {
	Banner []model.Article `json:"banner"`
	Up     []model.Article `json:"up"`
	List   []model.Article `json:"list"`
}

func (service CommonService) GetInitData(ctx *gin.Context) {
	if util.GetInt("basic.install") == 0 {
		util.ReturnMessage(ctx, false, "请完成初始化")
		return
	}
	ad := &model.ArticleData{
		Engine: service.Engine,
	}
	banner, _ := ad.GetList(true, false, "", -1, model.Official, 1, 20)
	up, _ := ad.GetList(false, true, "", -1, model.Official, 1, 20)
	list, _ := ad.GetList(true, true, "", -1, model.Official, 1, 20)
	data := &InitData{
		Banner: banner,
		Up:     up,
		List:   list,
	}
	util.ReturnData(ctx, true, data)
}

func (service CommonService) GetUserInfo(ctx *gin.Context) {
	check := util.GetString("github.bind")
	if check != "" {
		info, found := service.Cache.Get("user")
		if !found {
			var checkUser User
			err := json.Unmarshal([]byte(check), &checkUser)
			if err != nil {
				util.ReturnMessage(ctx, false, "获取博主信息失败")
			} else {
				checkUser.Id = ""
				checkUser.Version = util.Version
				service.Cache.Set("user", checkUser, cache.NoExpiration)
				util.ReturnData(ctx, true, checkUser)
			}
		} else {
			util.ReturnData(ctx, true, info)
		}
	} else {
		util.ReturnMessage(ctx, false, "系统尚未初始化")
	}
}

type UploadResponse struct {
	Msg  string             `json:"msg"`
	Code int                `json:"code"`
	Data UploadDataResponse `json:"data"`
}

type UploadDataResponse struct {
	ErrFiles []string          `json:"errFiles"`
	SuccMap  map[string]string `json:"succMap"`
}

func (service CommonService) Upload(ctx *gin.Context) {
	// 获取 multipart form
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		resp := UploadResponse{
			Msg:  "非法的上传格式",
			Code: 1,
			Data: UploadDataResponse{},
		}
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// 获取所有上传的文件
	files := ctx.Request.MultipartForm.File["files"]
	if len(files) == 0 {
		resp := UploadResponse{
			Msg:  "未能获取到文件",
			Code: 1,
			Data: UploadDataResponse{},
		}
		ctx.JSON(http.StatusOK, resp)
		return
	}

	var errFiles []string
	succMap := make(map[string]string)

	for _, file := range files {
		// 读取文件内容
		src, err := file.Open()
		if err != nil {
			errFiles = append(errFiles, file.Filename)
			continue
		}
		defer src.Close()

		// 将文件内容读入内存
		data, err := io.ReadAll(src)
		if err != nil {
			errFiles = append(errFiles, file.Filename)
			continue
		}

		hasher := md5.New()
		hasher.Write(data)
		filename := hex.EncodeToString(hasher.Sum(nil)) + ".jpg"

		// 检查文件是否已存在
		if _, err := os.Stat("./res/" + filename); !os.IsNotExist(err) {
			succMap[file.Filename] = "/res/" + filename
			continue
		}

		jpgData, err := util.Compression(data)
		if err != nil {
			errFiles = append(errFiles, file.Filename)
			continue
		}

		// 如果目录不存在，则创建
		if _, err := os.Stat("./res"); os.IsNotExist(err) {
			os.Mkdir("./res", 0755)
		}

		// 保存文件到 res 目录下
		save, err := os.Create("./res/" + filename)
		if err != nil {
			log.Println(err)
			errFiles = append(errFiles, file.Filename)
			continue
		}
		defer save.Close()

		_, err = save.Write(jpgData)
		if err != nil {
			errFiles = append(errFiles, file.Filename)
			continue
		}

		succMap[file.Filename] = "/res/" + filename
	}

	// 构建返回数据
	resp := UploadResponse{
		Msg:  "文件上传完成",
		Code: 0,
		Data: UploadDataResponse{
			ErrFiles: errFiles,
			SuccMap:  succMap,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
