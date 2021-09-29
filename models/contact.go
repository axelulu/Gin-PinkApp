package models

import "time"

type ContactAdd struct {
	SendId string `json:"send_id" form:"send_id"`
}

type Contact struct {
	UserId     int64     `json:"user_id" db:"user_id"`
	SendId     int64     `json:"send_id" db:"send_id"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type ContactDetail struct {
	UserId       int64       `json:"user_id" db:"user_id"`
	SendId       int64       `json:"send_id" db:"send_id"`
	SendUserMeta interface{} `json:"send_user_meta"`
	UpdateTime   time.Time   `json:"update_time" db:"update_time"`
}

type Message struct {
	Id         int64     `json:"id,omitempty" form:"id" db:"id"`                //消息ID
	UserId     int64     `json:"user_id,omitempty" form:"user_id" db:"user_id"` //谁发的
	Cmd        int       `json:"cmd,omitempty" form:"cmd" db:"cmd"`             //群聊还是私聊
	SendId     int64     `json:"send_id,omitempty" form:"send_id" db:"send_id"` //对端用户ID/群ID
	Media      int       `json:"media,omitempty" form:"media" db:"media"`       //消息按照什么样式展示
	Content    string    `json:"content,omitempty" form:"content" db:"content"` //消息的内容
	Pic        string    `json:"pic,omitempty" form:"pic" db:"pic"`             //预览图片
	Url        string    `json:"url,omitempty" form:"url" db:"url"`             //服务的URL
	Memo       string    `json:"memo,omitempty" form:"memo" db:"memo"`          //简单描述
	Amount     int       `json:"amount,omitempty" form:"amount" db:"amount"`    //其他和数字相关的
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
