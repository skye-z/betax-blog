package core

import (
	"github.com/gin-gonic/gin"
	"github.com/skye-z/betax-blog/util"
)

type SystemService struct {
}

func CreateSystemService() *SystemService {
	return &SystemService{}
}

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

func (service SystemService) GetConfig(ctx *gin.Context) {
	util.ReturnData(ctx, true, service.getConfig())
}

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

// 更新配置
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
	util.Set("basic.sslCert", config.SSLCert)
	util.Set("basic.sslKey", config.SSLKey)
	util.Set("ai.secret", config.AISecret)
	util.Set("ai.server", config.AIServer)
	util.Set("ai.setting", config.AISetting)
	util.Set("github.clientId", config.GithubClientId)
	util.Set("github.clientSecret", config.GithubClientSecret)
	util.Set("github.redirectUrl", config.GithubRedirectUrl)
	util.ReturnData(ctx, true, "初始化成功")
}
