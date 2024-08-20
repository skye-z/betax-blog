package core

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/betax-blog/util"
	githubreleases "github.com/skye-z/github-releases"
)

type SystemService struct {
	Version *githubreleases.Versioning
}

func CreateSystemService() *SystemService {
	ss := new(SystemService)
	ss.Version = &githubreleases.Versioning{
		Author: "skye-z",
		Store:  "betax-blog",
		Name:   "betax-blog",
		Cmd:    exec.Command("systemctl", "restart", "betax-blog"),
		Proxy:  "https://mirror.ghproxy.com/",
	}
	return ss
}

// 连通性测试
func (service SystemService) Ping(ctx *gin.Context) {
	util.ReturnMessage(ctx, true, "延迟测试")
}

type SystemConfig struct {
	Install            int    `json:"install"`            // basic.install
	SSLCert            string `json:"sslCert"`            // basic.sslCert
	SSLKey             string `json:"sslKey"`             // basic.sslKey
	AISecret           string `json:"aiSecret"`           // ai.secret
	AIServer           string `json:"aiSserver"`          // ai.server
	AIModel            string `json:"aiModel"`            // ai.model
	AISetting          string `json:"aiSetting"`          // ai.setting
	GithubClientId     string `json:"githubClientId"`     // github.clientId
	GithubClientSecret string `json:"githubClientSecret"` // github.clientSecret
	GithubRedirectUrl  string `json:"githubRedirectUrl"`  // github.redirectUrl
}

// 获取配置
func (service SystemService) GetConfig(ctx *gin.Context) {
	util.ReturnData(ctx, true, service.getConfig())
}

// 获取配置
func (service SystemService) getConfig() *SystemConfig {
	config := &SystemConfig{
		Install:            util.GetInt("basic.install"),
		SSLCert:            util.GetString("basic.sslCert"),
		SSLKey:             util.GetString("basic.sslKey"),
		AISecret:           util.GetString("ai.secret"),
		AIServer:           util.GetString("ai.server"),
		AIModel:            util.GetString("ai.model"),
		AISetting:          util.GetString("ai.setting"),
		GithubClientId:     util.GetString("github.clientId"),
		GithubClientSecret: util.GetString("github.clientSecret"),
		GithubRedirectUrl:  util.GetString("github.redirectUrl"),
	}
	return config
}

// 更新配置
func (service SystemService) UpdateConfig(ctx *gin.Context) {
	var config SystemConfig
	err := ctx.Bind(&config)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	old := service.getConfig()
	if old.SSLCert != config.SSLCert {
		util.Set("basic.sslCert", config.SSLCert)
	}
	if old.SSLKey != config.SSLKey {
		util.Set("basic.sslKey", config.SSLKey)
	}
	if old.AISecret != config.AISecret {
		util.Set("ai.secret", config.AISecret)
	}
	if old.AIServer != config.AIServer {
		util.Set("ai.server", config.AIServer)
	}
	if old.AISetting != config.AISetting {
		util.Set("ai.setting", config.AISetting)
	}
	if old.GithubClientId != config.GithubClientId {
		util.Set("github.clientId", config.GithubClientId)
	}
	if old.GithubClientSecret != config.GithubClientSecret {
		util.Set("github.clientSecret", config.GithubClientSecret)
	}
	if old.GithubRedirectUrl != config.GithubRedirectUrl {
		util.Set("github.redirectUrl", config.GithubRedirectUrl)
	}
	util.ReturnMessage(ctx, true, "")
}

// 初始化配置
func (service SystemService) Install(ctx *gin.Context) {
	if util.GetInt("basic.install") == 1 {
		util.ReturnData(ctx, true, "已初始化完成")
	}
	var config SystemConfig
	err := ctx.Bind(&config)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	util.Set("basic.install", 1)
	if config.SSLCert != "" {
		util.Set("basic.sslCert", config.SSLCert)
	}
	if config.SSLKey != "" {
		util.Set("basic.sslKey", config.SSLKey)
	}
	if config.AISecret != "" {
		util.Set("ai.secret", config.AISecret)
	}
	if config.AIServer != "" {
		util.Set("ai.server", config.AIServer)
	}
	if config.AISetting != "" {
		util.Set("ai.setting", config.AISetting)
	}
	if config.GithubClientId != "" {
		util.Set("github.clientId", config.GithubClientId)
	}
	if config.GithubClientSecret != "" {
		util.Set("github.clientSecret", config.GithubClientSecret)
	}
	if config.GithubRedirectUrl != "" {
		util.Set("github.redirectUrl", config.GithubRedirectUrl)
	}
	util.ReturnData(ctx, true, "初始化成功")
}

type FileInfo struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Modified int64  `json:"time"`
}

// 获取文件列表
func (service SystemService) GetFileList(ctx *gin.Context) {
	var files []FileInfo

	err := filepath.Walk("./res", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if !strings.HasSuffix(info.Name(), ".jpg") {
				return nil
			}
			files = append(files, FileInfo{
				Name:     info.Name(),
				Size:     info.Size(),
				Modified: info.ModTime().Unix() * 1000,
			})
		}
		return nil
	})

	if err != nil {
		util.ReturnMessage(ctx, false, "获取文件列表失败")
	} else {
		util.ReturnData(ctx, true, files)
	}
}

// 删除文件
func (service SystemService) RemoveFile(ctx *gin.Context) {
	name := ctx.PostForm("name")
	if len(name) != 36 {
		util.ReturnMessage(ctx, false, "文件不存在")
		return
	}
	err := os.Remove("./res/" + name)
	if err == nil {
		util.ReturnMessage(ctx, true, "删除成功")
	} else {
		log.Println(err)
		util.ReturnMessage(ctx, false, "文件删除失败")
	}
}

// 获取最新版本
func (service SystemService) GetNewVersion(ctx *gin.Context) {
	info := service.Version.GetLatestReleaseVersion()
	if info == nil {
		util.ReturnMessage(ctx, false, "获取版本信息失败")
	} else {
		util.ReturnData(ctx, true, info)
	}
}

// 更新版本
func (service SystemService) UpdateNewVersion(ctx *gin.Context) {
	state := service.Version.DownloadNewVersion()
	if state {
		util.ReturnMessage(ctx, true, "更新成功")
		service.Version.RestartWithSystemd()
	} else {
		util.ReturnMessage(ctx, false, "更新失败")
	}
}
