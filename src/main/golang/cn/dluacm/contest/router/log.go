package router

import (
	handler "acm-registration-system/handler/log"
	"acm-registration-system/middleware/jwt"
	"acm-registration-system/pkg/consts"
	"github.com/gin-gonic/gin"
)

func RegisterLogRouter(engin *gin.RouterGroup) {
	log := engin.Group("/log")
	log.Use(jwt.MustAuth(consts.Adminitrator))
	{
		log.GET("/:size/:num", handler.GetLogList)
		log.GET("/date/:start/:end/:size/:num", handler.GetLogListByDate)
		log.GET("/method/:methodId/:size/:num", handler.GetLogListByMethod)
		log.GET("/state/:stateId/:size/:num", handler.GetLogListByState)
	}
}
