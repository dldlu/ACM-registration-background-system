package router

import (
	handler "acm-registration-system/handler/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(engin *gin.RouterGroup) {
	user := engin.Group("/user")
	loginGroup := user.Group("/login")
	{
		loginGroup.POST("/adminitrator", handler.AdminitratorLogin)
	}
}
