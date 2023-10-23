package handler

import (
	"acm-registration-system/pkg/result"
	"acm-registration-system/pkg/types"
	service "acm-registration-system/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminitratorLogin
//
//	@Description: 管理员登录 参数: userLoginReq
//	@param c
func AdminitratorLogin(c *gin.Context) {
	var userLoginReq types.UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginReq); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res, err := service.GetUserService().AdminitratorLogin(userLoginReq.Number, userLoginReq.Password); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(res))
	}
}
