package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(500, gin.H{"msg": "未找到请求路由的处理函数"})
	}
}

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "方法不被允许"})
	}
}

// SkipperFunc 定义中间件跳过函数
type SkipperFunc func(*gin.Context) bool

// AllowPathPrefixSkipper 检查请求路径是否包含指定的前缀，如果包含则跳过
func AllowPathPrefixSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// AllowPathPrefixNoSkipper 检查请求路径是否包含指定的前缀，如果包含则不跳过
func AllowPathPrefixNoSkipper(prefixes ...string) SkipperFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

//AllowMethodAndPathPrefixSkipper 检查请求方法和路径是否包含指定的前缀，如果包含则跳过
func AllowMethodAndPathPrefixSkipper(prefixs ...string) SkipperFunc {
	return func(context *gin.Context) bool {
		path := JoinRouter(context.Request.Method, context.Request.URL.Path)
		pathlen := len(path)

		for _, p := range prefixs {
			if pl := len(p); pathlen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

//JoinRouter 拼接路由
func JoinRouter(method, path string) string {
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	return fmt.Sprintf("%s%s", strings.ToUpper(method), path)
}
