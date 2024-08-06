package core

import (
	"github.com/gin-gonic/gin"
	"github.com/skye-z/betax-blog/model"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

type TagService struct {
	Data *model.TagData
}

func CreateTagService(engine *xorm.Engine) *TagService {
	data := &model.TagData{
		Engine: engine,
	}
	return &TagService{
		Data: data,
	}
}

func (service TagService) GetList(ctx *gin.Context) {
	list, err := service.Data.GetSmallList()
	if err != nil {
		util.ReturnMessage(ctx, false, "获取标签列表失败")
	} else {
		util.ReturnData(ctx, true, list)
	}
}
