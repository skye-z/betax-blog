package main

import (
	"embed"
	"flag"

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
	// 定义一个命令行参数
	debug := flag.Bool("debug", false, "output debug logs")
	// 定义一个命令行参数
	port := flag.Int("port", 9800, "the port to listen on")
	// 解析命令行参数
	flag.Parse()
	// 初始化路由器
	router := core.BuildRouter(!*debug, *port, "0.0.0.0", "", "", engine, page)
	router.Run()
}
