package models

type Search struct {
	Type     string `json:"type" form:"type"`
	PostType string `json:"post_type" form:"post_type"`
	Word     string `json:"word" form:"word"`
	Page     int64  `json:"page" form:"page"`
	Size     int64  `json:"size" form:"size"`
}
