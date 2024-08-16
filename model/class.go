package model

import "xorm.io/xorm"

// 分类模型
type Class struct {
	Id    int64  `json:"id"`       // 分类编号
	Name  string `json:"name"`     // 分类名称
	Super int64  `json:"superior"` // 上级分类
}

type ClassData struct {
	Engine *xorm.Engine
}

// 获取分类列表
func (db ClassData) GetList() ([]Class, error) {
	var list []Class
	err := db.Engine.Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 新增分类
func (db ClassData) Add(class *Class) bool {
	_, err := db.Engine.Insert(class)
	return err == nil
}

// 编辑分类
func (db ClassData) Edit(class *Class) bool {
	if class.Id == 0 {
		return false
	}
	_, err := db.Engine.ID(class.Id).Update(class)
	return err == nil
}

// 删除分类
func (db ClassData) Del(id int64) bool {
	article := &Article{Class: -1}
	_, err := db.Engine.Where("class = ?", id).Update(article)
	if err != nil {
		return false
	}
	class := &Class{
		Id: id,
	}
	_, err = db.Engine.Delete(class)
	return err == nil
}
