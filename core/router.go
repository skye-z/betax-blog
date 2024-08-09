package core

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/skye-z/betax-blog/util"
	"xorm.io/xorm"
)

const MODEL_NAME = "Core"

type Router struct {
	// 路由对象
	Object *gin.Engine
	// 端口号
	port string
	// 监听地址
	host string
	// 证书地址
	cert string
	// 公钥地址
	key string
}

// 构建路由器
func BuildRouter(release bool, port int, host, cert, key string, engine *xorm.Engine, page embed.FS) *Router {
	if release {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = io.Discard
	router := &Router{
		Object: gin.Default(),
		cert:   cert,
		key:    key,
	}
	// 配置端口
	if port == 0 {
		if cert != "" && key != "" {
			router.port = "443"
		} else {
			router.port = "80"
		}
	} else {
		router.port = fmt.Sprint(port)
	}
	// 配置监听地址
	if host == "" {
		router.host = "0.0.0.0"
	} else {
		router.host = host
	}
	// 开启GZip压缩
	router.Object.Use(gzip.Gzip(gzip.BestSpeed))
	// 载入静态页面
	appPage, _ := fs.Sub(page, "page/app_dist")
	consolePage, _ := fs.Sub(page, "page/console_dist")
	router.Object.StaticFS("/app", http.FS(appPage))
	router.Object.StaticFS("/console", http.FS(consolePage))

	cs := CreateClassService(engine)
	ts := CreateTagService(engine)
	as := CreateArticleService(engine)
	ss := CreateSystemService()

	// 挂载鉴权路由
	addOAuth2Route(router.Object)
	// 挂载公共路由
	addPublicRoute(router.Object, cs, ts, as)
	// 挂载私有路由
	addPrivateRoute(router.Object, cs, ts, as, ss)
	return router
}

// 挂载公共路由
func addPublicRoute(router *gin.Engine, cs *ClassService, ts *TagService, as *ArticleService) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/app"
		router.HandleContext(ctx)
	})
	// 初始化接口(系统基础配置)
	// router.POST("/api/init", )
	// 初始化接口(导航栏、页脚栏、布局信息等)
	// router.GET("/api/init", )

	// 分类列表接口
	router.GET("/api/class", cs.GetList)

	// 标签列表接口
	router.GET("/api/tags", ts.GetList)

	// 文章列表接口(轮播文章列表、推荐文章列表、查询文章列表)
	router.POST("/api/article/list", as.GetList)
	// 文章数量接口
	router.POST("/api/article/number", as.GetNumber)
	// 文章搜索接口
	router.POST("/api/article/search", as.Search)
	// 文章详情接口
	router.GET("/api/article/details/:id", as.GetInfo)

	// 用户登陆接口
	// router.POST("/api/login", )
}

// 挂载私有路由
func addPrivateRoute(router *gin.Engine, cs *ClassService, ts *TagService, as *ArticleService, ss *SystemService) {
	private := router.Group("").Use(AuthHandler())
	{
		// 时间接口
		private.GET("/api/ping", ss.Ping)
		// 系统设置接口
		// private.POST("/api/setting", )
		// 垃圾清理接口
		// private.POST("/api/clean", )

		// 创建文章接口
		private.POST("/api/article/add", as.Add)
		// 编辑文章接口
		private.POST("/api/article/edit", as.Edit)
		// 发布文章接口
		private.POST("/api/article/publish", as.Publish)
		// 切换状态接口
		private.POST("/api/article/switch", as.SwitchState)
		// 文章AI摘要接口
		private.GET("/api/article/abstract", as.Abstract)
		// 删除文章接口
		private.POST("/api/article/remove", as.Remove)

		// 批量移动分类接口
		// private.POST("/api/article/move/class", )
		// 批量更新标签接口
		// private.POST("/api/article/update/tag", )

		// 新增分类接口
		private.POST("/api/class/add", cs.Add)
		// 编辑分类接口
		private.POST("/api/class/edit", cs.Edit)
		// 删除分类接口
		private.POST("/api/class/remove", cs.Remove)
	}
}

// 授权登陆路由
func addOAuth2Route(router *gin.Engine) {
	as := NewAuthService()
	if as != nil {
		router.GET("/login", as.Login)
		router.GET("/oauth2/callback", as.Callback)
	}
}

// 启动路由
func (r Router) Run() {
	util.OutLogf(MODEL_NAME, "starting the router from port "+r.port)
	// 启动服务
	go func() {
		if r.cert == "" {
			if err := r.Object.Run(":" + r.port); err != nil {
				util.OutErr(MODEL_NAME, "router startup failed: %v", err)
			}
		} else {
			if err := r.Object.RunTLS(":"+r.port, r.cert, r.key); err != nil {
				util.OutErr(MODEL_NAME, "router startup failed: %v", err)
			}
		}
	}()
	util.OutLog(MODEL_NAME, "router is ready")
	// 等待中断信号
	r.wait()
}

// 等待关闭
func (r Router) wait() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	util.OutLog(MODEL_NAME, "router is offline")

	// defer db.Close()

	util.OutLog(MODEL_NAME, "server is stopped")
}
