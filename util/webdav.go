package util

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const ModelName = "Sync"

// 同步
func Sync() {
	enable := GetBool("sync.enable")
	basePath := GetString("sync.path")
	if !enable || basePath == "" {
		return
	}
	token := getToken()
	// 运行模式 0无需同步, 1待拉取, 2待推送
	mode := 0
	// 获取云端同步标识
	content, err := downloadFile(token, basePath, ".synclog")
	if err != nil {
		if err.Error() == "404" {
			mode = 2
		} else {
			OutErr(ModelName, "unable to read cloud sync record")
			return
		}
	}
	// 云端存在同步信息
	if mode == 0 {
		// 解析云端同步标识
		cloud, err := strconv.ParseInt(string(content), 10, 64)
		if err != nil {
			OutErr(ModelName, "cloud sync record exception")
			return
		}
		if _, err := os.Stat("./.synclog"); os.IsExist(err) {
			// 本地有数据, 读取本地同步标识
			content, err := os.ReadFile("./.synclog")
			if err != nil {
				OutErr(ModelName, "unable to read local sync record")
				return
			}
			syncLog, err := strconv.ParseInt(string(content), 10, 64)
			if err != nil {
				OutErr(ModelName, "local sync record exception")
			}
			// 比对标识判断是否需要同步
			if cloud > syncLog {
				// 云端版本新于本地, 需要拉取
				mode = 1
			} else if cloud < syncLog {
				// 云端版本旧于本地, 需要推送
				mode = 2
			}
		} else {
			// 本地无数据, 需要拉取
			mode = 1
		}
	} else {
		// 云端没有数据, 需要推送
		mode = 2
	}
	// 无需同步, 直接退出
	if mode == 0 {
		return
	}
	// 需要同步, 拉取文件列表
	list, err := listFiles(token, basePath)
	if err != nil {
		log.Println(err)
		OutErr(ModelName, "failed to retrieve the list of cloud files")
		return
	}
	resList, err := listFiles(token, basePath+"/res")
	if mode == 1 {
		// 从云端拉取
		pullDB(token, basePath, list.Responses)
		if err != nil {
			subList := make([]Response, 0)
			pullRes(token, basePath, subList)
		} else {
			pullRes(token, basePath, resList.Responses)
		}
		pullComplete(token, basePath)
	} else if mode == 2 {
		// 从本地推送
		pushDB(token, basePath, list.Responses)
		if err != nil {
			subList := make([]Response, 0)
			pushRes(token, basePath, subList)
		} else {
			pushRes(token, basePath, resList.Responses)
		}
		pushComplete(token, basePath)
	}
}

// 上传文件
func uploadFile(token, path, filename string, content []byte) error {
	OutLog(ModelName, "upload "+filename)
	url := path + "/" + filename
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d", resp.StatusCode)
	}

	return nil
}

// 下载文件
func downloadFile(token, path, filename string) ([]byte, error) {
	OutLog(ModelName, "download "+filename)
	url := path + "/" + filename
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// 列出目录中的文件
func listFiles(token, path string) (*Multistatus, error) {
	req, err := http.NewRequest("PROPFIND", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Depth", "1")
	req.Header.Add("Authorization", "Basic "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusMultiStatus {
		return nil, fmt.Errorf("failed to list files, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var list Multistatus
	err = xml.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return &list, nil
}

// 创建文件夹
func mkdir(token, path, name string) error {
	OutLog(ModelName, "create folder")
	req, err := http.NewRequest("MKCOL", path+"/"+name, bytes.NewBufferString(""))
	if err != nil {
		return err
	}
	req.Header.Set("Depth", "0")
	req.Header.Add("Authorization", "Basic "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d", resp.StatusCode)
	}

	return nil
}

// 获取令牌
func getToken() string {
	auth := GetString("sync.username") + ":" + GetString("sync.password")
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// 拉取
func pullDB(token, path string, files []Response) {
	exist := false
	var cloud Response
	// 获取云端数据库信息
	for _, item := range files {
		if strings.HasSuffix(item.Href, "/local.store") {
			cloud = item
			exist = true
		}
	}
	// 检查本地数据库是否更新
	if exist {
		// 获取本地文件更新时间
		info, err := os.Stat("./local.store")
		if err != nil {
			return
		}
		check := checkCloudTime(info.ModTime(), cloud.PropStat.Prop.GetLastModified)
		// 云端是否更加新
		if check != 1 {
			return
		}
	}
	// 拉取
	data, err := downloadFile(token, path, "local.store")
	if err != nil {
		OutErr(ModelName, "unable to read local.store file")
		return
	}
	err = saveFile("./local.store", data)
	if err != nil {
		OutErr(ModelName, "unable to pull local.store file")
	}
}

// 拉取
func pullRes(token, path string, files []Response) {
	pullList := make([]string, 0)
	for _, item := range files {
		// 忽略目录
		if strings.HasSuffix(item.Href, "/") {
			continue
		}
		name := filepath.Base(item.Href)
		state, err := os.Stat("./res/" + name)
		if err != nil {
			log.Println(err)
			return
		}
		check := checkCloudTime(state.ModTime(), item.PropStat.Prop.GetLastModified)
		if check == 1 {
			pullList = append(pullList, name)
		}
	}

	localFiles, err := os.ReadDir("./res/")
	if err == nil {
		for _, x := range files {
			exist := false
			for _, y := range localFiles {
				if strings.HasSuffix(x.Href, "/"+y.Name()) {
					exist = true
					break
				}
			}
			if !exist {
				pullList = append(pullList, filepath.Base(x.Href))
			}
		}
	}

	// 拉取
	for _, name := range pullList {
		if strings.HasPrefix(name, ".") {
			continue
		}
		data, err := downloadFile(token, path, "./res/"+name)
		if err != nil {
			OutErr(ModelName, "unable to read "+name+" file")
			return
		}
		err = saveFile("./res/"+name, data)
		if err != nil {
			log.Println(err)
			OutErr(ModelName, "unable to pull "+name+" file")
		}
	}
}

// 拉取完成
func pullComplete(token, path string) {
	var now string
	// 拉取云端时间
	data, err := downloadFile(token, path, ".synclog")
	if err != nil {
		now = strconv.FormatInt(time.Now().Unix(), 10)
		OutErr(ModelName, "unable to read .synclog file")
	} else {
		now = string(data)
	}
	// 写入本地标识文件
	file, err := os.OpenFile("./.synclog", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		OutErr(ModelName, "unable to operate .synclog file")
	}
	defer file.Close()
	_, err = file.WriteString(now)
	if err != nil {
		OutErr(ModelName, "unable to write .synclog file")
	}
}

// 推送数据库
func pushDB(token, path string, files []Response) {
	res := false
	exist := false
	var cloud Response
	for _, item := range files {
		if strings.HasSuffix(item.Href, "/local.store") {
			cloud = item
			exist = true
		} else if strings.HasSuffix(item.Href, "/res/") {
			res = true
		}
	}
	if !res {
		err := mkdir(token, path, "res")
		if err != nil {
			OutErr(ModelName, "Unable to create 'res' folder")
		}
	}
	// 检查本地数据库是否更新
	if exist {
		// 获取本地文件更新时间
		info, err := os.Stat("./local.store")
		if err != nil {
			return
		}
		check := checkLocalTime(cloud.PropStat.Prop.GetLastModified, info.ModTime())
		// 本地是否更加新
		if check != 1 {
			return
		}
	}
	// 推送
	data, err := readFile("./local.store")
	if err != nil {
		OutErr(ModelName, "unable to read local.store file")
		return
	}
	err = uploadFile(token, path, "local.store", data)
	if err != nil {
		OutErr(ModelName, "unable to upload local.store file")
	}
}

// 推送资源文件
func pushRes(token, path string, files []Response) {
	pushList := make([]string, 0)
	for _, item := range files {
		// 忽略目录
		if strings.HasSuffix(item.Href, "/") {
			continue
		}
		name := filepath.Base(item.Href)
		state, err := os.Stat("./res/" + name)
		if err != nil {
			log.Println(err)
			return
		}
		check := checkLocalTime(item.PropStat.Prop.GetLastModified, state.ModTime())
		if check == 1 {
			pushList = append(pushList, name)
		}
	}

	localFiles, err := os.ReadDir("./res/")
	if err == nil {
		for _, x := range localFiles {
			exist := false
			for _, y := range files {
				if strings.HasSuffix(y.Href, "/"+x.Name()) {
					exist = true
					break
				}
			}
			if !exist {
				pushList = append(pushList, x.Name())
			}
		}
	}

	// 推送
	for _, name := range pushList {
		if strings.HasPrefix(name, ".") {
			continue
		}
		data, err := readFile("./res/" + name)
		if err != nil {
			OutErr(ModelName, "unable to read "+name+" file")
			return
		}
		err = uploadFile(token, path+"/res", name, data)
		if err != nil {
			log.Println(err)
			OutErr(ModelName, "unable to upload "+name+" file")
		}
	}
}

// 推送完成
func pushComplete(token, path string) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	// 写入本地标识文件
	file, err := os.OpenFile("./.synclog", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		OutErr(ModelName, "unable to operate .synclog file")
	}
	defer file.Close()
	_, err = file.WriteString(now)
	if err != nil {
		OutErr(ModelName, "unable to write .synclog file")
	}
	// 上传标识文件
	err = uploadFile(token, path, ".synclog", []byte(now))
	if err != nil {
		OutErr(ModelName, "unable to upload .synclog file")
	}
}

// 读取本地文件
func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 保存本地文件
func saveFile(path string, date []byte) error {
	// 写入本地标识文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		OutErr(ModelName, "unable to operate .synclog file")
		return err
	}
	defer file.Close()
	_, err = file.Write(date)
	if err != nil {
		OutErr(ModelName, "unable to write .synclog file")
		return err
	}
	return nil
}

// 校验时间
// 1	本地新
// -1	云端新
func checkLocalTime(ct string, lt time.Time) int {
	if ct == "" {
		return 1
	}
	layout := "Mon, _2 Jan 2006 15:04:05 GMT"
	cloudTime, err := time.Parse(layout, ct)
	if err != nil {
		return 0
	}
	// 本地是否更加新
	if lt.After(cloudTime) {
		return 1
	} else if lt.Before(cloudTime) {
		return -1
	}
	return 0
}
func checkCloudTime(ct time.Time, lt string) int {
	if lt == "" {
		return 1
	}
	layout := "Mon, _2 Jan 2006 15:04:05 GMT"
	cloudTime, err := time.Parse(layout, lt)
	if err != nil {
		return 0
	}
	// 云端是否更加新
	if ct.After(cloudTime) {
		return 1
	} else if ct.Before(cloudTime) {
		return -1
	}
	return 0
}

type Resource struct {
	CreationDate     string       `xml:"creationdate"`
	DisplayName      string       `xml:"displayname"`
	GetLastModified  string       `xml:"getlastmodified"`
	GetContentLength int          `xml:"getcontentlength"`
	GetContentType   string       `xml:"getcontenttype"`
	GetETag          string       `xml:"getetag"`
	ResourceType     ResourceType `xml:"resourcetype"`
}

type ResourceType struct {
	IsCollection bool `xml:"collection"`
}

type PropStat struct {
	Prop   Resource `xml:"prop"`
	Status string   `xml:"status"`
}

type Response struct {
	Href     string   `xml:"href"`
	PropStat PropStat `xml:"propstat"`
}

type Multistatus struct {
	XMLName   xml.Name   `xml:"multistatus"`
	Responses []Response `xml:"response"`
}
