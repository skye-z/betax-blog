/*
字符串工具

BetaX Harbor
Copyright © 2024 SkyeZhang <skai-zhang@hotmail.com>
*/

package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateRandomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func CalculateMD5(input string) string {
	data := []byte(input)
	hash := md5.Sum(data)
	hashInHex := hex.EncodeToString(hash[:])
	return hashInHex
}

func GetPostBool(ctx *gin.Context, name string, defaultValue bool) bool {
	cache := ctx.PostForm(name)
	if cache == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(cache)
	if err != nil {
		return defaultValue
	}
	return value
}

func GetPostInt(ctx *gin.Context, name string, defaultValue int) int {
	cache := ctx.PostForm(name)
	if cache == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(cache)
	if err != nil {
		return defaultValue
	}
	return value
}
