package util

import (
	"github.com/skye-z/betax-blog/model"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func InitDB() *xorm.Engine {
	OutLogf("Data", "load engine")
	engine, err := xorm.NewEngine("sqlite", "./local.store")
	if err != nil {
		panic(err)
	}
	return engine
}

func InitDBTable(engine *xorm.Engine) {
	OutLogf("Data", "load table")
	err := engine.Sync2(new(model.Class))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Tag))
	if err != nil {
		panic(err)
	}
	err = engine.Sync2(new(model.Article))
	if err != nil {
		panic(err)
	}
}
