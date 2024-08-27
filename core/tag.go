package core

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/skye-z/betax-blog/model"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

type TagService struct {
	Data  *model.TagData
	Cache *cache.Cache
}

func CreateTagService(engine *xorm.Engine, cache *cache.Cache) *TagService {
	data := &model.TagData{
		Engine: engine,
	}
	return &TagService{
		Data:  data,
		Cache: cache,
	}
}

func (service TagService) GetList(ctx *gin.Context) {
	if util.GetInt("basic.install") == 0 {
		util.ReturnMessage(ctx, false, "请完成初始化")
		return
	}

	info, found := service.Cache.Get("tags")
	if !found {
		list, err := service.Data.GetSmallList()
		if err != nil {
			util.ReturnMessage(ctx, false, "获取标签列表失败")
		} else {
			service.Cache.Set("tags", list, cache.NoExpiration)
			util.ReturnData(ctx, true, list)
		}
	} else {
		util.ReturnData(ctx, true, info)
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
		service.Cache.Delete("tags")
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
		service.Cache.Delete("tags")
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
		service.Cache.Delete("tags")
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "标签删除失败")
	}
}
