package models

import "time"

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username     string `json:"username" form:"username" binding:"required"`
	Email        string `json:"email" form:"email" binding:"required"`
	ValidateCode string `json:"validate_code" form:"validate_code" binding:"required"`
	Password     string `json:"password" form:"password" binding:"required"`
	RePassword   string `json:"re_password" form:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserForgetPwd struct {
	Email         string `json:"email" form:"email" binding:"required"`
	ValidateCode  string `json:"validate_code" form:"validate_code" binding:"required"`
	NewPassword   string `json:"password" form:"password" binding:"required"`
	ReNewPassword string `json:"re_password" form:"re_password" binding:"required,eqfield=NewPassword"`
}

type UserPost struct {
	UserId   int64  `json:"user_id" form:"user_id"`
	PostType string `json:"post_type" form:"post_type"`
	Page     int64  `json:"page" form:"page"`
	Size     int64  `json:"size" form:"size"`
}

type UserUpdate struct {
	Slug  string `json:"slug" form:"slug" binding:"required,oneof=avatar background username gender birth descr"`
	Value string `json:"value" form:"value" binding:"required"`
}

type UserPasswordUpdate struct {
	OldPassword   string `json:"old_password" form:"old_password"`
	Email         string `json:"email" form:"email" binding:"required"`
	ValidateCode  string `json:"validate_code" form:"validate_code" binding:"required"`
	NewPassword   string `json:"new_password" form:"new_password" binding:"required"`
	ReNewPassword string `json:"re_new_password" form:"re_new_password" binding:"required,eqfield=NewPassword"`
}

type UserEmailUpdate struct {
	Email        string `json:"new_email" form:"new_email" binding:"required"`
	ValidateCode string `json:"validate_code" form:"validate_code" binding:"required"`
}

type User struct {
	UserID     int64     `json:"user_id" db:"user_id"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Background string    `json:"background" db:"background"`
	Username   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Email      string    `json:"email" db:"email"`
	Fans       int64     `json:"fans" db:"fans"`
	Follows    int64     `json:"follows" db:"follows"`
	Coin       int64     `json:"coin" db:"coin"`
	IsVip      time.Time `json:"is_vip" db:"is_vip"`
	Birth      time.Time `json:"birth" db:"birth"`
	Gender     int64     `json:"gender" db:"gender"`
	Active     int       `json:"active"`
}

type UserMeta struct {
	UserID     int64     `json:"user_id" db:"user_id"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Background string    `json:"background" db:"background"`
	Username   string    `json:"username" db:"username"`
	Descr      string    `json:"descr" db:"descr"`
	Email      string    `json:"email" db:"email"`
	Fans       int64     `json:"fans" db:"fans"`
	Follows    int64     `json:"follows" db:"follows"`
	Coin       int64     `json:"coin" db:"coin"`
	Phone      int64     `json:"phone" db:"phone"`
	Exp        int64     `json:"exp" db:"exp"`
	Gender     int64     `json:"gender" db:"gender"`
	Birth      time.Time `json:"birth" db:"birth"`
	IsVip      time.Time `json:"is_vip" db:"is_vip"`
	Active     int64     `json:"active"`
}

type FansMeta struct {
	IsFollow bool
	UserMeta *UserMeta
}
