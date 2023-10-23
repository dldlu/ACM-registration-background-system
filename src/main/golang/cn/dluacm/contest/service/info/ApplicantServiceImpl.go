package service

import (
	"acm-registration-system/dal/db"
	"acm-registration-system/dal/model"
	"acm-registration-system/pkg/errno"
	"acm-registration-system/pkg/types"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type ApplicantServiceImpl struct {
}

func (c ApplicantServiceImpl) GetApplicants(Size int, Num int) (*types.PageResp, *errno.Errno) {
	applicants, err := db.Information.Where(db.Information.ID.Gt(0)).Limit(Size).Offset(Size * (Num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	count, err2 := db.Information.Where(db.Information.ID.Gt(0)).Count()
	if err2 != nil {
		logrus.Error(err2)
		return nil, errno.NewErrno(errno.DbErrorCode, err2.Error())
	}
	return &types.PageResp{ItemTotal: count, PageTotal: count/int64(Size) + 1, Array: applicants}, nil
}

func (c ApplicantServiceImpl) GetApplicant(id int64) (*model.Information, *errno.Errno) {
	applicants, err := db.Information.Where(db.Information.ID.Eq(id)).Take()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return applicants, nil
}

func (c ApplicantServiceImpl) PostApplicantByYear(Year int, HalfYear int, Size int, Num int) (*types.PageResp, *errno.Errno) {
	var applicants []*model.Information
	var err error
	var timeStart time.Time
	var timeEnd time.Time
	if HalfYear == -1 {
		timeStart, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprint(Year, "-01-01 01:01:01"))
		timeEnd, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprint(Year, "-12-31 23:59:59"))
	} else if HalfYear == 1 {
		timeStart, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprint(Year, "-01-01 01:01:01"))
		timeEnd, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprint(Year, "-06-30 23:59:59"))
	} else if HalfYear == 2 {
		timeStart, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprint(Year, "-06-30 23:59:59"))
		timeEnd, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprint(Year, "-12-31 23:59:59"))
	}
	applicants, err = db.Information.Where(db.Information.SignTime.Between(timeStart, timeEnd)).Limit(Size).Offset(Size * (Num - 1)).Find()
	if err != nil {
		logrus.Error(err)
		return nil, errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	count, err2 := db.Information.Where(db.Information.SignTime.Between(timeStart, timeEnd)).Limit(Size).Offset(Size * (Num - 1)).Count()
	if err2 != nil {
		logrus.Error(err2)
		return nil, errno.NewErrno(errno.DbErrorCode, err2.Error())
	}
	return &types.PageResp{ItemTotal: count, PageTotal: count/int64(Size) + 1, Array: applicants}, nil
}

func (c ApplicantServiceImpl) PostApplicant(applicant *model.Information) *errno.Errno {
	err := db.Information.Create(&model.Information{ID: applicant.ID, StudentNumber: applicant.StudentNumber,
		StudentName: applicant.StudentName, MajorClass: applicant.MajorClass, Sex: applicant.Sex, Qq: applicant.Qq,
		Phone: applicant.Phone, Email: applicant.Email, Year: applicant.Year, LastIP: applicant.LastIP, SignTime: applicant.SignTime,
		Status: applicant.Status})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ApplicantServiceImpl) PutAapplicant(applicant *model.Information) *errno.Errno {
	fmt.Println(applicant.SignTime)
	_, err := db.Information.Where(db.Information.ID.Eq(applicant.ID)).Updates(&model.Information{ID: applicant.ID,
		StudentNumber: applicant.StudentNumber, StudentName: applicant.StudentName, MajorClass: applicant.MajorClass,
		Sex: applicant.Sex, Qq: applicant.Qq, Phone: applicant.Phone, Email: applicant.Email, Year: applicant.Year,
		LastIP: applicant.LastIP, SignTime: applicant.SignTime, Status: applicant.Status})
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}

func (c ApplicantServiceImpl) DeleteAapplicants(applicantIds []int64) *errno.Errno {
	_, err := db.Information.Where(db.Information.ID.In(applicantIds...)).Delete()
	if err != nil {
		logrus.Error(err)
		return errno.NewErrno(errno.DbErrorCode, err.Error())
	}
	return nil
}
