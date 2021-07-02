package middleware

import "github.com/gin-gonic/gin"

/**
casbin拦截器功能
 */

//CasbinMiddleware casbin中间件
func CasbinMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
		//用户ID
/*		uid,isExit := c.Get(common.USER_ID_Key)
		if !isExit {
			common.ResFailCode(c,"token 无效3",50008)
			return
		}*/
	}
}
