package handler

import (
	"acm-registration-system/dal/model"
	"acm-registration-system/pkg/result"
	"acm-registration-system/pkg/types"
	service "acm-registration-system/service/info"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetApplicants
//
//	@Description: 分页查询所有报名人员信息 参数:page
//	@param c
func GetApplicants(c *gin.Context) {
	var page types.PageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	if applicants, err := service.GetInfoService().GetApplicants(page.Size, page.Num); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(applicants))
	}
}

// GetApplicant
//
//	@Description: 根据id查询单个报名人员信息 参数: IdRequest
//	@param c
func GetApplicant(c *gin.Context) {
	var idRequest types.IdRequest
	if err := c.ShouldBindUri(&idRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	if applicant, err := service.GetInfoService().GetApplicant(idRequest.Id); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(applicant))
	}
}

// PostApplicantByYear
//
//	@Description: 根据年份报名人员信息 参数: page
//	@param c
func PostApplicantByYear(c *gin.Context) {
	var page types.YearPageRequest
	if err := c.ShouldBindUri(&page); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	if applicants, err := service.GetInfoService().PostApplicantByYear(page.Year, page.HalfYear, page.Size, page.Num); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.NewOkResult(applicants))
	}
}

// PostApplicant
//
//	@Description: 添加报名人员信息 参数: applicantRequest
//	@param c
func PostAapplicant(c *gin.Context) {
	var applicantRequest model.Information
	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	if err := service.GetInfoService().PostApplicant(&applicantRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// PutAapplicant
//
//	@Description: 修改单个报名人员信息 参数: applicantRequest
//	@param c
func PutAapplicant(c *gin.Context) {
	var applicantRequest model.Information
	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	if err := service.GetInfoService().PutAapplicant(&applicantRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// DeleteAapplicants
//
//	@Description: 批量删除报名人员信息 参数: applicantRequest
//	@param c
func DeleteAapplicants(c *gin.Context) {
	var applicantRequest any
	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.JSON(http.StatusBadRequest, result.NewFailedResult(err.Error()))
		return
	}
	str := fmt.Sprint(applicantRequest)
	// 提取方括号内的数字字符串
	idsStr := strings.Split(str, "ids:")[1]
	idsStr = strings.Trim(idsStr, "[]")
	// 将数字字符串分割为单个数字
	ids := make([]int64, 0)
	for _, idStr := range strings.Split(idsStr, " ") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		ids = append(ids, int64(id))
	}
	if err := service.GetInfoService().DeleteAapplicants(ids); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}
