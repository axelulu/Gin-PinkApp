package models

import "time"

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

type UserPost struct {
	UserId   int64  `json:"user_id" form:"user_id"`
	PostType string `json:"post_type" form:"post_type"`
	Page     int64  `json:"page" form:"page"`
	Size     int64  `json:"size" form:"size"`
}

type User struct {
	UserID     int64     `json:"user_id" db:"user_id"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Background string    `json:"background" db:"background"`
	Username   string    `json:"username" db:"username"`
	Password   string    `json:"password" db:"password"`
	Fans       int64     `json:"fans" db:"fans"`
	Follows    int64     `json:"follows" db:"follows"`
	Coin       int64     `json:"coin" db:"coin"`
	IsVip      time.Time `json:"is_vip" db:"is_vip"`
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
	Active     int       `json:"active"`
}

type FansMeta struct {
	IsFollow bool
	UserMeta *UserMeta
}
