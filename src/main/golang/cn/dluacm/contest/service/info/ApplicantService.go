package service

import (
	"acm-registration-system/dal/model"
	"acm-registration-system/pkg/errno"
	"acm-registration-system/pkg/types"
)

type ApplicantService interface {
	// GetApplicants
	// @Description: 分页查询所有报名人员信息
	// @param Size int, Num int
	// @return *types.PageResp,*errno.Errno
	GetApplicants(Size int, Num int) (*types.PageResp, *errno.Errno)

	// GetApplicant
	// @Description: 根据id查询单个报名人员信息
	// @param id int64
	// @return *model.Information, *errno.Errno
	GetApplicant(id int64) (*model.Information, *errno.Errno)

	// PostApplicantByYear
	// @Description: 根据年份分页查询报名人员信息
	// @param Year int,Size int, Num int
	// @return *types.PageResp,*errno.Errno
	PostApplicantByYear(Year int, HalfYear int, Size int, Num int) (*types.PageResp, *errno.Errno)

	// PostApplicant
	// @Description: 添加报名人员信息
	// @param *model.Information
	// @return *errno.Errno
	PostApplicant(*model.Information) *errno.Errno

	// PutAapplicant
	// @Description: 根据id修改单个报名人员信息
	// @param *model.Information
	// @return *errno.Errno
	PutAapplicant(*model.Information) *errno.Errno

	// DeleteAapplicants
	// @Description: 批量删除报名人员信息
	// @param applicantIds []int64
	// @return *errno.Errno
	DeleteAapplicants(applicantIds []int64) *errno.Errno
}
