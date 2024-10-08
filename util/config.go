/*
配置工具

BetaX Blog
Copyright © 2024 SkyeZhang <skai-zhang@hotmail.com>
*/

package util

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/spf13/viper"
)

const Version = "0.2.0"

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("ini")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefault()
		} else {
			OutErr("Config", "read failed: %v", err)
		}
	}
}

func Reload() {
	viper.ReadInConfig()
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
	viper.WriteConfig()
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func createDefault() {
	// 安装状态
	viper.SetDefault("basic.install", "0")
	viper.SetDefault("basic.sslCert", "")
	viper.SetDefault("basic.sslKey", "")
	// AI
	viper.SetDefault("ai.secret", "")
	viper.SetDefault("ai.server", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions")
	viper.SetDefault("ai.model", "qwen2-1.5b-instruct")
	viper.SetDefault("ai.setting", "你是个摘要大师, 擅长总结Markdown和HTML中的内容, 你写摘要内容要点明确且语言简练, 你写的摘要可以控制在300字以内")
	// OAuth2
	viper.SetDefault("github.clientId", "")
	viper.SetDefault("github.clientSecret", "")
	viper.SetDefault("github.redirectUrl", "")
	viper.SetDefault("github.bind", "")
	// 同步
	viper.SetDefault("sync.enable", "false")
	viper.SetDefault("sync.path", "")
	viper.SetDefault("sync.username", "")
	viper.SetDefault("sync.password", "")
	// 令牌密钥
	secret, err := generateSecret()
	if err != nil {
		panic(err)
	}
	viper.SetDefault("token.secret", secret)
	// 令牌有效期/小时
	viper.SetDefault("token.exp", 24)
	viper.SafeWriteConfig()
}

func generateSecret() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
