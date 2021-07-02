package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go_casbin/routers/api/v1"
	"go_casbin/routers/casbin/enforcer"
	"go_casbin/routers/jwt"
	"go_casbin/utils"
)

func RegisterRouterSys(app *gin.RouterGroup)  {
/*	menu:=sys.Menu{}
	app.GET("/menu/list", menu.List)
	app.GET("/menu/detail", menu.Detail)
	app.GET("/menu/allmenu", menu.AllMenu)
	app.GET("/menu/menubuttonlist", menu.MenuButtonList)
	app.POST("/menu/delete", menu.Delete)
	app.POST("/menu/update", menu.Update)
	app.POST("/menu/create", menu.Create)
	user := sys.User{}
	app.GET("/user/info", user.Info)
	app.POST("/user/login", user.Login)
	app.POST("/user/logout", user.Logout)
	app.POST("/user/editpwd", user.EditPwd)
	admins := sys.Admins{}
	app.GET("/admins/list", admins.List)
	app.GET("/admins/detail", admins.Detail)
	app.GET("/admins/adminsroleidlist", admins.AdminsRoleIDList)
	app.POST("/admins/delete", admins.Delete)
	app.POST("/admins/update", admins.Update)
	app.POST("/admins/create", admins.Create)
	app.POST("/admins/setrole", admins.SetRole)
	role := sys.Role{}
	app.GET("/role/list", role.List)
	app.GET("/role/detail", role.Detail)
	app.GET("/role/rolemenuidlist", role.RoleMenuIDList)
	app.GET("/role/allrole", role.AllRole)
	app.POST("/role/delete", role.Delete)
	app.POST("/role/update", role.Update)
	app.POST("/role/create", role.Create)
	app.POST("/role/setrole", role.SetRole)*/
	group1 := app.Group("")
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
}
