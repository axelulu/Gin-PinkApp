package models

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserID   int64  `json:"user_id" db:"user_id"`
	Avatar   string `json:"avatar" db:"avatar"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Fans     int64  `json:"fans" db:"fans"`
}
