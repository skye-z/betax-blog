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
	if util.GetInt("basic.install") == 0 {
		util.ReturnMessage(ctx, false, "请完成初始化")
		return
	}
	list, err := service.Data.GetSmallList()
	if err != nil {
		util.ReturnMessage(ctx, false, "获取标签列表失败")
	} else {
		util.ReturnData(ctx, true, list)
	}
}

// 添加标签
func (service TagService) Add(ctx *gin.Context) {
	var tag model.Tag
	err := ctx.Bind(&tag)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	if service.Data.Add(&tag) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "创建标签失败")
	}
}

// 编辑标签
func (service TagService) Edit(ctx *gin.Context) {
	var tag model.Tag
	err := ctx.Bind(&tag)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	if service.Data.Edit(&tag) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "编辑标签失败")
	}
}

// 删除标签
func (service TagService) Remove(ctx *gin.Context) {
	tagId := util.GetPostInt64(ctx, "id", 0)
	if tagId <= 0 {
		util.ReturnMessage(ctx, false, "标签编号无效")
		return
	}
	if service.Data.Del(tagId) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "标签删除失败")
	}
}
