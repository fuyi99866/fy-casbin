package manageweb

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_casbin/conf"
	"go_casbin/docs"
	"go_casbin/middleware"
	"go_casbin/models"
	"go_casbin/pkg/logger"
	"go_casbin/routers"
	"net/http"
	"time"
)

//运行
func Run() {
	//1、加载配置文件
	config := flag.String("c", "./app.ini", "config file path")
	//解析配置文件
	flag.Parse()
	conf.Init(*config) //根据配置文件初始化配置
	//2、初始化日志
	logger.InitLog("info", "go_casbin.log")

	//3、初始化数据库
	initDB()
	//4、初始化casbin
	//5、初始化web服务
	initWeb()
	logger.Info("--------服务启动-------")

}

//初始化数据库
func initDB() {
	models.Init()
}

//初始化网站
func initWeb() {
	gin.SetMode(gin.DebugMode) //调试模式
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())
	//崩溃恢复
	app.Use(middleware.RecoveryMiddleware())
	app.LoadHTMLGlob( "dist/*.html")
	app.Static("/static","dist/static")
	app.Static("/resource", "resource")
	app.StaticFile("/favicon.ico","dist/favicon.ico")

	//注册路由
	routers.RegisterRouter(app)
	go initHTTPServer(app)
	//go RunServer()
}

// InitHTTPServer 初始化http服务
func initHTTPServer(handler http.Handler) {
	srv := &http.Server{
		Addr:         ":"+fmt.Sprintf(":%d", conf.ServerSetting.HttpPort),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}


func RunServer() {
	if conf.Swag != nil {
		docs.SwaggerInfo.Host = conf.Swag.Host
		docs.SwaggerInfo.BasePath = conf.ServerSetting.BasePath
		scheme := "http"
		if conf.ServerSetting.HTTPS {
			scheme = "https"
		}
		logger.Info(fmt.Sprintf("-----服务启动,可以打开  %s://%s%s/swagger/index.html 查看详细接口------", scheme, conf.Swag.Host, conf.ServerSetting.BasePath, ))
	}

	routers.Run()
}
