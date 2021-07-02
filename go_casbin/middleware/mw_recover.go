package middleware

import (
	"github.com/gin-gonic/gin"
	"go_casbin/pkg/logger"
)

//RecoveryMiddleware 崩溃恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		logger.Info("崩溃恢复中间件")
	}
}