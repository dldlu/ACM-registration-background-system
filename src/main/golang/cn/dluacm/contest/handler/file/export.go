package handler

import (
	"acm-registration-system/pkg/result"
	service "acm-registration-system/service/file"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// PostExport
//
//	@Description: 输出下载文件 参数:
//	@param c
func PostExport(c *gin.Context) {
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
	if err := service.GetFileService().PostExport(ids); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result.Ok())
	}
}

// GetExport
//
//	@Description: 输出下载文件 参数:
//	@param c
func GetExport(c *gin.Context) {
	fileName := "Applicant.xlsx"
	filePath := path.Join("./", fileName)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filePath)
	os.Remove(filePath)
}
