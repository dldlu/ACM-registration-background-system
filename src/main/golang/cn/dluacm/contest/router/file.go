package router

import (
	handler "acm-registration-system/handler/file"
	"acm-registration-system/middleware/jwt"
	"acm-registration-system/pkg/consts"

	"github.com/gin-gonic/gin"
)

func RegisterFileRouter(engin *gin.RouterGroup) {
	file := engin.Group("/file")
	{
		file.POST("/export", jwt.MustAuth(consts.Adminitrator), handler.PostExport)
		file.GET("/export", jwt.MustAuth(consts.Adminitrator), handler.GetExport)
	}
}
