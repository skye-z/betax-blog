package core

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skye-z/betax-blog/model"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

type ArticleService struct {
	Data *model.ArticleData
}

func CreateArticleService(engine *xorm.Engine) *ArticleService {
	data := &model.ArticleData{
		Engine: engine,
	}
	return &ArticleService{
		Data: data,
	}
}

// 获取文章列表
func (service ArticleService) GetList(ctx *gin.Context) {
	// 是否轮播
	isBanner := util.GetPostBool(ctx, "isBanner", false)
	// 是否置顶
	isUp := util.GetPostBool(ctx, "isUp", false)
	// 筛选关键词
	keyword := ctx.PostForm("keyword")
	// 文章分类
	class := util.GetPostInt(ctx, "class", -1)
	// 文章状态
	state := util.GetPostInt(ctx, "state", 2)
	// 页码
	page := util.GetPostInt(ctx, "page", 1)
	// 数量
	num := util.GetPostInt(ctx, "number", 20)
	// 获取列表
	list, err := service.Data.GetList(isBanner, isUp, keyword, class, state, page, num)
	if err != nil {
		util.ReturnMessage(ctx, false, "获取文章列表失败")
	} else {
		util.ReturnData(ctx, true, list)
	}
}

// 获取文章数量
func (service ArticleService) GetNumber(ctx *gin.Context) {
	// 筛选关键词
	keyword := ctx.PostForm("keyword")
	// 文章状态
	state := util.GetPostInt(ctx, "state", 0)
	// 获取数量
	number, err := service.Data.GetNumber(keyword, state)
	if err != nil {
		util.ReturnMessage(ctx, false, "获取文章数量失败")
	} else {
		util.ReturnData(ctx, true, number)
	}
}

// 搜索文章
func (service ArticleService) Search(ctx *gin.Context) {
	// 筛选关键词
	keyword := ctx.PostForm("keyword")
	// 页码
	page := util.GetPostInt(ctx, "page", 1)
	// 数量
	num := util.GetPostInt(ctx, "number", 20)
	// 搜索符合的文章
	list, err := service.Data.Search(keyword, page, num)
	if err != nil {
		util.ReturnMessage(ctx, false, "搜索文章失败")
	} else {
		util.ReturnData(ctx, true, list)
	}
}

// 获取文章详情
func (service ArticleService) GetInfo(ctx *gin.Context) {
	cache := ctx.Param("id")
	articleId, err := strconv.ParseInt(cache, 10, 64)
	if err != nil {
		util.ReturnMessage(ctx, false, "文章编号无效")
		return
	}
	info := service.Data.Get(articleId)
	if info == nil {
		util.ReturnMessage(ctx, false, "获取文章详情失败")
	} else {
		util.ReturnData(ctx, true, info)
	}
}

// 添加文章
func (service ArticleService) Add(ctx *gin.Context) {
	var article model.Article
	err := ctx.Bind(&article)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	// 写入创建时间
	article.CreationTime = time.Now().Unix() * 1000
	if service.Data.Add(&article) {
		util.ReturnData(ctx, true, article.Id)
	} else {
		util.ReturnMessage(ctx, false, "创建文章失败")
	}
}

// 编辑文章
func (service ArticleService) Edit(ctx *gin.Context) {
	var article model.Article
	err := ctx.Bind(&article)
	if err != nil {
		util.ReturnMessage(ctx, false, "表单内容无效")
		return
	}
	// 写入更新时间
	article.LastUpdateTime = time.Now().Unix() * 1000
	if service.Data.Edit(&article) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "编辑文章失败")
	}
}

// 发布文章
func (service ArticleService) Publish(ctx *gin.Context) {
	articleId := util.GetPostInt64(ctx, "id", 0)
	if articleId == 0 {
		util.ReturnMessage(ctx, false, "文章编号无效")
		return
	}
	article := service.Data.Get(articleId)
	article.LastUpdateTime = time.Now().Unix() * 1000
	article.ReleaseTime = time.Now().Unix() * 1000
	article.State = model.Official
	if service.Data.Edit(article) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "发布文章失败")
	}
}

// 切换状态
func (service ArticleService) SwitchState(ctx *gin.Context) {
	articleId := util.GetPostInt64(ctx, "id", 0)
	if articleId == 0 {
		util.ReturnMessage(ctx, false, "文章编号无效")
		return
	}
	state := util.GetPostInt(ctx, "state", 2)
	article := service.Data.Get(articleId)
	article.LastUpdateTime = time.Now().Unix() * 1000
	article.State = state
	if service.Data.Edit(article) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "文章状态更新失败")
	}
}

// 删除文章
func (service ArticleService) Remove(ctx *gin.Context) {
	articleId := util.GetPostInt64(ctx, "id", 0)
	if articleId == 0 {
		util.ReturnMessage(ctx, false, "文章编号无效")
		return
	}
	if service.Data.Del(articleId) {
		util.ReturnData(ctx, true, nil)
	} else {
		util.ReturnMessage(ctx, false, "文章删除失败")
	}
}

// 文章摘要
func (service ArticleService) Abstract(ctx *gin.Context) {
	cache := ctx.Query("id")
	articleId, err := strconv.ParseInt(cache, 10, 64)
	if err != nil {
		util.ReturnMessage(ctx, false, "文章编号无效")
		return
	}
	article := service.Data.Get(articleId)
	if len(article.Content) < 20 {
		util.ReturnMessage(ctx, false, "文章内容太少啦")
		return
	}
	util.ReturnData(ctx, true, util.OpenAISingleRound(article.Content))
}

// 获取随机列表
func (service ArticleService) GetAnyList(ctx *gin.Context) {
	// 获取列表
	list, err := service.Data.GetAnyList()
	if err != nil {
		util.ReturnMessage(ctx, false, "获取随机列表失败")
	} else {
		util.ReturnData(ctx, true, list)
	}
}
