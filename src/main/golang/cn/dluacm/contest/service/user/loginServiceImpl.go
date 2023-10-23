package service

import (
	"acm-registration-system/dal/db"
	"acm-registration-system/pkg/errno"
	"acm-registration-system/pkg/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type LoginServiceImpl struct {
}

// 默认账号密码
// 12345678
// Dlu@12345678
func (c LoginServiceImpl) AdminitratorLogin(number, password string) (map[string]interface{}, *errno.Errno) {
	// fmt.Println(encodePassword(password))
	res := make(map[string]interface{}, 1)
	adminitrator, err := db.AdminInfo.Where(db.AdminInfo.Username.Eq(number)).
		Select(db.AdminInfo.ID, db.AdminInfo.Username, db.AdminInfo.Password, db.AdminInfo.RoleID).Take()
	if err != nil {
		return nil, errno.NewErrno(errno.DbErrorCode, errno.DefaultErrMsg)
	}
	if checkPassword(password, adminitrator.Password) {
		token := utils.GenerateToken(adminitrator.ID, adminitrator.RoleID, number)
		res["token"] = token
		return res, nil
	} else {
		return nil, errno.NewErrno(errno.DefaultErrCode, errno.DefaultErrMsg)
	}
}

// CheckPassword 检验密码
func checkPassword(passwordInput, passwordSave string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordSave), []byte(passwordInput))
	fmt.Println(err)
	return err == nil
}

// // encodePassword 加密
// func encodePassword(password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), consts.PassWordCost)
// 	logrus.Info(err)
// 	return string(bytes)
// }
