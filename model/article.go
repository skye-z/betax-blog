package model

import (
	"xorm.io/xorm"
)

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
	Id             int64   `json:"id"`                // 编号
	Weight         int     `json:"weight"`            // 权重
	IsBanner       bool    `json:"isBanner"`          // 横幅轮播
	IsUp           bool    `json:"isUp"`              // 置顶
	Type           int     `json:"type"`              // 类型(见上方枚举)
	Title          string  `json:"title"`             // 标题
	Abstract       string  `json:"abstract"`          // 摘要
	Class          int64   `json:"class"`             // 分类
	TagIds         []int64 `json:"tagIds" xorm:"-"`   // 传入标签
	Tags           *[]Tag  `json:"tags" xorm:"-"`     // 传出标签
	Content        string  `json:"content,omitempty"` // 内容
	State          int     `json:"state"`             // 状态(见上方枚举)
	CreationTime   int64   `json:"creationTime"`      // 创建时间
	LastUpdateTime int64   `json:"lastUpdateTime"`    // 上次更新时间
	ReleaseTime    int64   `json:"releaseTime"`       // 发布时间
}

type ArticleData struct {
	Engine *xorm.Engine
}

func (a *Article) SetTags(newTags []Tag) {
	a.Tags = &newTags
}

// 搜索文章
func (db ArticleData) Search(keyword string, page int, num int) ([]Article, error) {
	var list []Article
	err := db.Engine.Where("state = 2 AND title LIKE ? OR abstract LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Omit("content").Desc("weight", "release_time", "creation_time").Limit(page*num, (page-1)*num).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取文章列表
func (db ArticleData) GetList(isBanner, isUp bool, keyword string, class, state, page, num int) ([]Article, error) {
	var list []Article
	var cache *xorm.Session
	if isBanner && isUp {
		cache = db.Engine.Where("1=1")
	} else {
		cache = db.Engine.Where("is_banner = ? AND is_up = ?", isBanner, isUp)
	}
	if class >= 0 {
		cache.And("class = ?", class)
	}
	if state != 0 {
		cache.And("state = ?", state)
	}
	if keyword != "" {
		cache.And("title LIKE ? OR abstract LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	err := cache.Omit("content").Desc("weight", "creation_time", "release_time").Limit(num, (page-1)*num).Find(&list)
	if err != nil {
		return nil, err
	}
	tagData := &TagData{
		Engine: db.Engine,
	}
	for i, item := range list {
		tags, _ := tagData.GetList(item.Id)
		(&list[i]).SetTags(tags)
	}
	return list, nil
}

// 获取文章列表
func (db ArticleData) GetListByTag(tag, page, num int) ([]Article, error) {
	var list []Article
	err := db.Engine.Table("article").Join("LEFT", "tag_connect", "article.id = tag_connect.article").
		Omit("content").Where("tag_connect.tag = ?", tag).Desc("weight", "creation_time", "release_time").Limit(num, (page-1)*num).Find(&list)
	if err != nil {
		return nil, err
	}
	tagData := &TagData{
		Engine: db.Engine,
	}
	for i, item := range list {
		tags, _ := tagData.GetList(item.Id)
		(&list[i]).SetTags(tags)
	}
	return list, nil
}

func (db ArticleData) GetAnyList() ([]Article, error) {
	var list []Article
	err := db.Engine.SQL("SELECT id,type,title,abstract,class,creation_time,last_update_time,release_time,is_up FROM article ORDER BY RANDOM() LIMIT 10").Find(&list)
	if err != nil {
		return nil, err
	}
	tagData := &TagData{
		Engine: db.Engine,
	}
	for i, item := range list {
		tags, _ := tagData.GetList(item.Id)
		(&list[i]).SetTags(tags)
	}
	return list, nil
}

// 获取文章数量
func (db ArticleData) GetNumber(keyword string, state int) (int64, error) {
	var article Article
	if keyword == "" && state == 0 {
		return db.Engine.Count(article)
	} else {
		if keyword != "" && state != 0 {
			return db.Engine.Where("title LIKE ? OR abstract LIKE ? AND state = ?", "%"+keyword+"%", "%"+keyword+"%", state).Count(article)
		} else if keyword == "" {
			return db.Engine.Where("state = ?", state).Count(article)
		}
		return db.Engine.Where("title LIKE ? OR abstract LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Count(article)
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

	tagData := &TagData{
		Engine: db.Engine,
	}
	tags, _ := tagData.GetList(article.Id)
	article.SetTags(tags)
	return article
}

// 新增文章
func (db ArticleData) Add(article *Article) bool {
	_, err := db.Engine.Insert(article)
	tagData := &TagData{
		Engine: db.Engine,
	}
	tagData.UpdateConnect(article.Id, &article.TagIds)
	return err == nil
}

// 编辑文章
func (db ArticleData) Edit(article *Article) bool {
	if article.Id == 0 {
		return false
	}
	_, err := db.Engine.ID(article.Id).Update(article)
	if err != nil {
		return false
	}
	tagData := &TagData{
		Engine: db.Engine,
	}
	tagData.UpdateConnect(article.Id, &article.TagIds)
	return err == nil
}

// 编辑文章
func (db ArticleData) EditMeta(article *Article) bool {
	if article.Id == 0 {
		return false
	}
	_, err := db.Engine.ID(article.Id).Cols("is_up", "is_banner", "last_update_time").Update(article)
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
