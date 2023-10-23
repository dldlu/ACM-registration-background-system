package router

import (
	handler "acm-registration-system/handler/info"
	"acm-registration-system/middleware/jwt"
	"acm-registration-system/pkg/consts"

	"github.com/gin-gonic/gin"
)

func RegisterInfoRouter(engin *gin.RouterGroup) {
	info := engin.Group("/info")
	{
		info.GET("/applicants/:size/:num", jwt.MustAuth(consts.Adminitrator), handler.GetApplicants)
		info.GET("/applicant/:id", jwt.MustAuth(consts.Adminitrator), handler.GetApplicant)
		info.POST("/applicant/:year/:halfYear/:size/:num", jwt.MustAuth(consts.Adminitrator), handler.PostApplicantByYear)
		info.POST("/applicant", jwt.MustAuth(consts.Adminitrator), handler.PostAapplicant)
		info.PUT("/applicant", jwt.MustAuth(consts.Adminitrator), handler.PutAapplicant)
		info.DELETE("/applicant", jwt.MustAuth(consts.Adminitrator), handler.DeleteAapplicants)
	}
}
