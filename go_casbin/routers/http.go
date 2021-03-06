package routers

import (

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	v1 "go_casbin/routers/api/v1"
	"go_casbin/routers/casbin/enforcer"
	"go_casbin/routers/jwt"
	"go_casbin/utils"
)

//swag init --generalInfo .\routers\http.go

// @title code server
// @version 0.0.1
// @description Go 学习综合demo
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run() {
	//1.创建路由
	r := gin.Default()
	//2、绑定路由规则，执行的函数
	//gin.Context,封装了request和response

	//r.Use(enforcer.Interceptor(enforcer.EnforcerTool()))


	//将访问路由到swagger的HTML页面
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // API 注释
	r.POST("/auth", v1.Auth) // token鉴权

	//websocket实现聊天室
	//r.GET("/ws", websocket.NotifySocket)

	//group1 := r.Group("api/v1")
	group1 := r.Group("")
	group1.Use(jwt.JWT()) //token 验证
	group1.Use(enforcer.Interceptor(enforcer.EnforcerTool()))  //拦截器进行访问控制
	{
		group1.GET("/alive",v1.TokenAlive)
		user := group1.Group("user")
		{
			user.GET("/test", utils.Response_test)//测试回复
			user.GET("/:name",v1.GetApiParam)
			user.GET("",v1.GetUsers)
			user.POST("",v1.AddUser)
			user.DELETE("delete",v1.DeleteUser)
			user.POST("update",v1.UpdateUser)
			user.GET("get",v1.GetOneUser)
		}

		policy := group1.Group("policy")
		{
			policy.POST("",v1.AddPolicy)
			policy.DELETE("",v1.DeletePolicy)
			policy.GET("",v1.GetPolicy)
		}

		authority := group1.Group("authority")
		{
			authority.POST("add",v1.CreateAuthority)
			authority.POST("update",v1.UpdateAuthority)
			authority.POST("set",v1.SetAuthority)
			authority.DELETE("delete",v1.DeleteAuthority)

		}
	}


	//3、监听端口，默认是8080
	//Run("里面不指定端口就默认为8080")

	r.Run(":8082")
}
