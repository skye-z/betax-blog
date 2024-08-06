package model

import "xorm.io/xorm"

// 状态枚举
const (
	Draft    = iota + 1 // 草稿
	Official            // 公开
	Private             // 私密
	Remove              // 删除
)

// 类型枚举
const (
	Markdown = iota + 1 // MK
	HTML                // 网页
	Table               // 表格
	Graph               // 图表
)

// 文章模型
type Article struct {
	Id             int64  `json:"id"`             // 编号
	Weight         int    `json:"weight"`         // 权重
	IsBanner       bool   `json:"isBanner"`       // 横幅轮播
	IsUp           bool   `json:"isUp"`           // 置顶
	Type           int    `json:"type"`           // 类型(见上方枚举)
	Title          string `json:"title"`          // 标题
	Abstract       string `json:"abstract"`       // 摘要
	Class          int64  `json:"class"`          // 分类
	Content        string `json:"content"`        // 内容
	State          int    `json:"state"`          // 状态(见上方枚举)
	CreationTime   int64  `json:"creationTime"`   // 创建时间
	LastUpdateTime int64  `json:"lastUpdateTime"` // 上次更新时间
	ReleaseTime    int64  `json:"releaseTime"`    // 发布时间
}

type ArticleData struct {
	Engine *xorm.Engine
}

// 搜索文章
func (db ArticleData) Search(keyword string, page int, num int) ([]Article, error) {
	var list []Article
	err := db.Engine.Where("title LIKE ? OR abstract LIKE ?", keyword, keyword).Desc("weight", "releaseTime", "creationTime").Limit(page*num, (page-1)*num).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取文章列表
func (db ArticleData) GetList(isBanner, isUp bool, state, page, num int) ([]Article, error) {
	var list []Article
	err := db.Engine.Where("(isBanner = ? OR isUp = ?) AND state = ?", isBanner, isUp, state).Desc("weight", "releaseTime", "creationTime").Limit(page*num, (page-1)*num).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取文章数量
func (db ArticleData) GetNumber(keyword string) (int64, error) {
	var article Article
	if keyword == "" {
		return db.Engine.Count(article)
	} else {
		return db.Engine.Where("title LIKE ? OR abstract LIKE ?", keyword, keyword).Count(article)
	}
}

// 获取文章详情
func (db ArticleData) Get(id int64) *Article {
	article := &Article{
		Id: id,
	}
	has, _ := db.Engine.Get(article)
	if !has {
		return nil
	}
	return article
}

// 新增文章
func (db ArticleData) Add(article *Article) bool {
	_, err := db.Engine.Insert(article)
	return err == nil
}

// 编辑文章
func (db ArticleData) Edit(article *Article) bool {
	if article.Id == 0 {
		return false
	}
	_, err := db.Engine.ID(article.Id).Update(article)
	return err == nil
}

// 逻辑删除文章
func (db ArticleData) Del(id int64) bool {
	article := db.Get(id)
	if article == nil {
		return false
	}
	article.State = Remove
	return db.Edit(article)
}
