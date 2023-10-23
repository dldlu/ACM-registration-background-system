package service

import "acm-registration-system/pkg/errno"

type LoginService interface {
	// AdminitratorLogin
	// @Description: 管理员登陆
	// @param number 管理员账号
	// @param password 管理员密码
	// @return map[string]interface{} 响应
	// @return *errno.Errno 异常
	AdminitratorLogin(number, password string) (map[string]interface{}, *errno.Errno)
}
