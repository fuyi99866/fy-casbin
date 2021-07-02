package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_casbin/models"
	"go_casbin/pkg/logger"
	"go_casbin/routers/casbin/enforcer"
	"go_casbin/utils"
	"net/http"
)

// @Summary   增加访问权限
// @Tags   访问权限
// @Accept json
// @Produce  json
// @Param   body  body   models.UserPolicy   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {object} utils.Response
// @Router /policy  [POST]
// @Security ApiKeyAuth
func AddPolicy(c *gin.Context) {
	logger.Info("1增加Policy")
	appG := utils.Gin{C: c}
	var reqInfo models.UserPolicy
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		logger.Error("AddPolicy param error")
		appG.Response(http.StatusBadRequest, utils.INVALID_PARAMS, err.Error())
		return
	}

	e := enforcer.EnforcerTool()
	fmt.Println("增加Policy")
	if ok := e.AddPolicy(reqInfo.Username,reqInfo.URL,reqInfo.Type); !ok {
		logger.Info("Policy已经存在")
		appG.Response(http.StatusInternalServerError, utils.ERROR, nil)
	} else {
		logger.Info("增加成功")
		appG.Response(http.StatusOK, utils.SUCCESS, nil)
	}
}

// @Summary   删除访问权限
// @Tags   访问权限
// @Accept json
// @Produce  json
// @Param   body  body   models.UserPolicy   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {object} utils.Response
// @Router /policy  [DELETE]
// @Security ApiKeyAuth
func DeletePolicy(c *gin.Context) {
	logger.Info("删除Policy")
	appG := utils.Gin{C: c}
	var reqInfo models.UserPolicy
	err := c.ShouldBindJSON(&reqInfo)
	if err != nil {
		logger.Error("AddPolicy param error")
		appG.Response(http.StatusBadRequest, utils.INVALID_PARAMS, err.Error())
		return
	}

	e := enforcer.EnforcerTool()
	if ok := e.RemovePolicy(reqInfo.Username,reqInfo.URL,reqInfo.Type); !ok {
		logger.Info("Policy不存在")
		appG.Response(http.StatusInternalServerError, utils.ERROR, nil)
	} else {
		logger.Info("删除成功")
		appG.Response(http.StatusOK, utils.SUCCESS, nil)
	}
}

// @Summary   获取权限列表
// @Tags   访问权限
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {object} utils.Response
// @Router /policy  [GET]
// @Security ApiKeyAuth
func GetPolicy(c *gin.Context) {
	logger.Info("1查看Policy")
	appG := utils.Gin{C: c}
	e := enforcer.EnforcerTool()

	list := e.GetPolicy()
	for _, vlist := range list {
		for _, v := range vlist {
			logger.Info("value: ", v)
		}
	}

	appG.Response(http.StatusOK, utils.SUCCESS, list)
}
