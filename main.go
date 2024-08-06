package main

import (
	"embed"

	"github.com/skye-z/betax-blog/core"
	"github.com/skye-z/betax-blog/util"
)

//go:embed page/app_dist/* page/console_dist/*
var page embed.FS

func main() {
	// 初始化系统配置
	util.InitConfig()
	// 初始化数据库
	engine := util.InitDB()
	go util.InitDBTable(engine)
	// 初始化路由器
	router := core.BuildRouter(false, 9800, "0.0.0.0", "", "", engine, page)
	router.Run()
}
