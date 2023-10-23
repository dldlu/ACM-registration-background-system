package service

import "sync"

var userService *UserService
var userServiceOnce sync.Once

type UserService struct {
	LoginService
}

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userService = &UserService{}
		loginService := new(LoginServiceImpl)

		userService.LoginService = loginService
	})
	return userService
}
