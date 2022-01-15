package models

import (
	"time"
)

type Home struct {
	*PostCategoryList
	CSize int64 `json:"cSize" form:"cSize"`
}

type PostListByIds struct {
	PostIds string `json:"postIds" form:"postIds"`
	Page    int64  `json:"page" form:"page"`
	Size    int64  `json:"size" form:"size"`
}

type PostCategoryList struct {
	CategorySlug int64  `json:"category_id" form:"category_id"`
	Page         int64  `json:"page" form:"page"`
	Size         int64  `json:"size" form:"size"`
	Sort         string `json:"sort" form:"sort" binding:"oneof=rand update_time view likes reply"`
}

type PostAuthorList struct {
	AuthorId int64 `json:"author_id" form:"author_id"`
	Page     int64 `json:"page" form:"page"`
	Size     int64 `json:"size" form:"size"`
}

type PostRankingList struct {
	RankingSlug string `json:"ranking" form:"ranking"`
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
}

type PostDynamicList struct {
	DynamicSlug string `json:"dynamic" form:"dynamic"`
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
}

type Post struct {
	PostId       int64     `json:"post_id" db:"post_id"`
	AuthorId     int64     `json:"author_id" db:"author_id"`
	PostType     string    `json:"post_type" db:"post_type"`
	CategorySlug int64     `json:"category_id" db:"category_id"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	Reply        int64     `json:"reply" db:"reply"`
	Favorite     int64     `json:"favorite" db:"favorite"`
	Likes        int64     `json:"likes" db:"likes"`
	UnLikes      int64     `json:"un_likes" db:"un_likes"`
	Coin         int64     `json:"coin" db:"coin"`
	Share        int64     `json:"share" db:"share"`
	View         int64     `json:"view" db:"view"`
	Cover        string    `json:"cover" db:"cover"`
	Video        string    `json:"video" db:"video"`
	Download     string    `json:"download" db:"download"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}

type PostPublish struct {
	PostType     string `json:"post_type" form:"post_type" binding:"oneof=post video dynamic collection"`
	CategorySlug int64  `json:"category_id" form:"category_id"`
	Title        string `json:"title" form:"title"`
	Content      string `json:"content" form:"content"`
	Cover        string `json:"cover" form:"cover"`
	Download     string `json:"download" form:"download"`
	Video        string `json:"video" form:"video"`
}

type PostDetail struct {
	Owner interface{} `json:"owner"`
	*Post             // 帖子信息
}
