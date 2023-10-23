package service

import (
	"acm-registration-system/dal/db"
	"acm-registration-system/pkg/errno"
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExportServiceImpl struct {
}

func (c ExportServiceImpl) PostExport(applicantIds []int64) *errno.Errno {
	// 查询数据库
	rows, err := db.Information.Where(db.Information.ID.In(applicantIds...)).Find()
	if err != nil {
		return errno.NewErrno(errno.DbErrorCode, errno.DefaultErrMsg)
	}

	// 创建Excel文件
	file := excelize.NewFile()
	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	//根据数据建立表格
	rowsCount := len(rows)
	file.SetSheetRow(sheetName, "A1", &[]interface{}{"序号", "学号", "姓名", "专业班级", "性别", "QQ", "电话", "邮箱"})
	for i := 0; i < rowsCount; i++ {
		file.SetSheetRow(sheetName, fmt.Sprintf("A%d", i+2), &[]interface{}{i + 1, rows[i].StudentNumber,
			*(rows[i].StudentName), *(rows[i].MajorClass), rows[i].Sex, rows[i].Qq, rows[i].Phone, rows[i].Email})
	}
	//根据指定路径保存文件
	if err := file.SaveAs("Applicant.xlsx"); err != nil {
		return errno.NewErrno(errno.DbErrorCode, errno.DefaultErrMsg)
	}
	return nil
}
