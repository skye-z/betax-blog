package model

import (
	"xorm.io/xorm"
)

// 标签模型
type Tag struct {
	Id   int64  `json:"id"`   // 标签编号
	Name string `json:"name"` // 标签名称
}

// 标签连接模型
type TagConnect struct {
	Id      int64 `json:"id"`      // 连接编号
	Article int64 `json:"article"` // 文章
	Tag     int64 `json:"tag"`     // 标签
}

type TagData struct {
	Engine *xorm.Engine
}

// 获取少量分类列表
func (db TagData) GetSmallList() ([]Tag, error) {
	var list []Tag
	err := db.Engine.Limit(20, 0).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 获取分类列表
func (db TagData) GetList(article int64) ([]Tag, error) {
	var list []Tag
	var tagIds []int64
	var connect []TagConnect
	err := db.Engine.Where("article = ?", article).Find(&connect)
	if err != nil {
		return nil, err
	}
	for _, item := range connect {
		tagIds = append(tagIds, item.Tag)
	}
	err = db.Engine.In("id", tagIds).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 新增标签
func (db TagData) Add(tag *Tag) bool {
	_, err := db.Engine.Insert(tag)
	return err == nil
}

// 编辑标签
func (db TagData) Edit(tag *Tag) bool {
	if tag.Id == 0 {
		return false
	}
	_, err := db.Engine.ID(tag.Id).Update(tag)
	return err == nil
}

// 删除标签
func (db TagData) Del(id int64) bool {
	// 删除标签连接
	connect := &TagConnect{
		Tag: id,
	}
	_, err := db.Engine.Delete(connect)
	if err != nil {
		return false
	}
	// 删除标签本体
	tag := &Tag{
		Id: id,
	}
	_, err = db.Engine.Delete(tag)
	return err == nil
}

// 新增标签连接
func (db TagData) AddConnect(connect *TagConnect) bool {
	_, err := db.Engine.Insert(connect)
	return err == nil
}

// 删除标签连接
func (db TagData) DelConnect(article, tag int64) bool {
	connect := &TagConnect{
		Article: article,
		Tag:     tag,
	}
	_, err := db.Engine.Delete(connect)
	return err == nil
}

// 删除标签连接
func (db TagData) UpdateConnect(article int64, ids *[]int64) bool {
	var connect []TagConnect
	err := db.Engine.Where("article = ?", article).Find(&connect)
	if err != nil {
		return false
	}
	// 遍历现有标签
	for _, sup := range connect {
		var exist bool
		for _, sub := range *ids {
			if sup.Tag == sub {
				exist = true
				break
			}
		}
		// 判断传入标签是否存在
		if !exist {
			db.DelConnect(article, sup.Tag)
		}
	}
	// 遍历传入标签
	for _, sub := range *ids {
		var exist bool
		for _, sup := range connect {
			if sup.Tag == sub {
				exist = true
				break
			}
		}
		// 判断现有标签是否存在
		if !exist {
			obj := &TagConnect{
				Article: article,
				Tag:     sub,
			}
			db.AddConnect(obj)
		}
	}
	return true
}
