package jwt

import (
	"context"
	"acm-registration-system/pkg/errno"
	"acm-registration-system/pkg/result"
	"acm-registration-system/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var key int

type UserInfo struct {
	ID     int64
	Number string
	RoleId int64
}

// newContext
// @Description: 创建context储存数据
// @param ctx
// @param claims
// @return context.Context
func newContext(ctx context.Context, claims *UserInfo) context.Context {
	return context.WithValue(ctx, key, claims)
}

// GetUserInfo
// @Description: 获取context保存的数据
// @param ctx
// @return *UserInfo
func GetUserInfo(ctx context.Context) *UserInfo {
	u, _ := ctx.Value(key).(*UserInfo)
	return u
}

// MustAuth
// @Description: 必须为此权限
// @auth sinre 2023-05-12 22:16:05
// @param roleId
// @return gin.HandlerFunc
func MustAuth(roleId int64) gin.HandlerFunc {
	return func(context *gin.Context) {
		logrus.Info("==============开始认证===========")
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, result.NewResult(errno.DefaultErrCode, errno.DefaultErrMsg, nil))
		} else {
			auths := strings.Fields(auth)
			if len(auths) != 2 {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.NewResult(errno.DefaultErrCode, errno.DefaultErrMsg, nil))
			} else {
				auth = auths[1]
				claim, err := utils.ParseToken(auth)
				if err != nil {
					context.Abort()
					context.JSON(http.StatusUnauthorized, result.NewResult(errno.DefaultErrCode, errno.DefaultErrMsg, nil))
				} else if roleId != claim.RoleId {
					context.Abort()
					context.JSON(http.StatusUnauthorized, result.NewResult(errno.DefaultErrCode, errno.DefaultErrMsg, nil))
				} else {
					context.Request = context.Request.WithContext(newContext(context, &UserInfo{ID: claim.UserId, Number: claim.Number, RoleId: claim.RoleId}))
					logrus.Info("==============认证成功===========")
					context.Next()
				}
			}
		}
	}
}
