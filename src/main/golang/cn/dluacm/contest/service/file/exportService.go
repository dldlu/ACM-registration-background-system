package service

import (
	"acm-registration-system/pkg/errno"
)

type ExportService interface {
	// PostExport
	// @Description: 输出下载文件
	// @param
	// @return *errno.Errno
	PostExport(applicantIds []int64) *errno.Errno
}
