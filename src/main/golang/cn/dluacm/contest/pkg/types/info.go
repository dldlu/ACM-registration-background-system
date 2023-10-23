package types

import "time"

type ApplicantRequest struct {
	ID            int64     `json:"id"`             // 自增id
	StudentNumber string    `json:"student_number"` // 学号
	StudentName   *string   `json:"student_name"`   // 姓名
	MajorClass    *string   `json:"major_class"`    // 专业班级
	Sex           string    `json:"sex"`            // 性别
	Qq            string    `json:"qq"`             // qq号
	Phone         string    `json:"phone"`          // 手机号
	Email         string    `json:"email"`          // 邮箱
	Year          int64     `json:"year"`           // 报名年份
	LastIP        string    `json:"last_ip"`        // 提交ip
	SignTime      time.Time `json:"sign_time"`      // 提交时间
	Status        int8      `json:"status"`         // 状态位
}
