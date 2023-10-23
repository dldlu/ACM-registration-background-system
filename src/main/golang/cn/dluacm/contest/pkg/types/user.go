package types

// UserLoginRequest
// @Description: 用户登陆请求
type UserLoginRequest struct {
	Number   string `json:"number" binding:"required,max=16,min=8"`
	Password string `json:"password" binding:"required,max=16,min=8"`
}

