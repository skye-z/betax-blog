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
