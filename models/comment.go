package models

import "time"

type Comment struct {
	PostId     int64     `json:"post_id" db:"post_id"`
	UserId     int64     `json:"user_id" db:"user_id"`
	Content    string    `json:"content" db:"content"`
	Type       string    `json:"type" db:"type"`
	Parent     int64     `json:"parent" db:"parent"`
	LikeNum    int64     `json:"like_num" db:"like_num"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type CommentCreate struct {
	PostId  string `json:"post_id" form:"post_id"`
	Content string `json:"content" form:"content"`
	Type    string `json:"type" form:"type"`
	Parent  string `json:"parent" form:"parent"`
}

type CommentList struct {
	PostId string `json:"post_id" form:"post_id"`
	Page   int64  `json:"page" form:"page"`
	Size   int64  `json:"size" form:"size"`
}

type CommentDetail struct {
	Owner interface{} `json:"owner"`
	*Comment
}
