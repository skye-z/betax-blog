package core

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/skye-z/betax-blog/model"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

type ClassService struct {
	Data  *model.ClassData
	Cache *cache.Cache
}

func CreateClassService(engine *xorm.Engine, cache *cache.Cache) *ClassService {
	data := &model.ClassData{
		Engine: engine,
	}
	return &ClassService{
		Data:  data,
		Cache: cache,
	}
}

// 获取分类列表
func (service ClassService) GetList(ctx *gin.Context) {
	if util.GetInt("basic.install") == 0 {
		util.ReturnMessage(ctx, false, "请完成初始化")
		return
	}

	info, found := service.Cache.Get("class")
	if !found {
		list, err := service.Data.GetList()
		if err != nil {
			util.ReturnMessage(ctx, false, "获取分类列表失败")
		} else {
			service.Cache.Set("class", list, cache.NoExpiration)
			util.ReturnData(ctx, true, list)
		}
	} else {
		util.ReturnData(ctx, true, info)
	}
}

// 添加分类
func (service ClassService) Add(ctx *gin.Context) {
	var class model.Class
	err := ctx.Bind(&class)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	if service.Data.Add(&class) {
		service.Cache.Delete("class")
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "创建分类失败")
	}
}

// 编辑分类
func (service ClassService) Edit(ctx *gin.Context) {
	var class model.Class
	err := ctx.Bind(&class)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	if service.Data.Edit(&class) {
		service.Cache.Delete("class")
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "编辑分类失败")
	}
}

// 删除分类
func (service ClassService) Remove(ctx *gin.Context) {
	classId := util.GetPostInt64(ctx, "id", 0)
	if classId <= 0 {
		util.ReturnMessage(ctx, false, "分类编号无效")
		return
	}
	if service.Data.Del(classId) {
		service.Cache.Delete("class")
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "分类删除失败")
	}
}
