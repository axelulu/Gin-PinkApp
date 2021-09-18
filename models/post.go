package models

import "time"

type Home struct {
	*PostCategoryList
	CSize int64 `json:"cSize" form:"cSize"`
}

type PostCategoryList struct {
	CategorySlug string `json:"category_slug" form:"category_slug"`
	Page         int64  `json:"page" form:"page"`
	Size         int64  `json:"size" form:"size"`
}

type PostAuthorList struct {
	AuthorId int64 `json:"author_id" form:"author_id"`
	Page     int64 `json:"page" form:"page"`
	Size     int64 `json:"size" form:"size"`
}

type PostRankingList struct {
	RankingSlug string `json:"ranking" form:"rinking"`
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
}

type Post struct {
	PostId       int64     `json:"post_id" db:"post_id"`
	AuthorId     int64     `json:"author_id" db:"author_id"`
	PostType     string    `json:"post_type" db:"post_type"`
	CategorySlug string    `json:"category_slug" db:"category_slug"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	Reply        int64     `json:"reply" db:"reply"`
	Favorite     int64     `json:"favorite" db:"favorite"`
	Likes        int64     `json:"likes" db:"likes"`
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
	PostType     string `json:"post_type" db:"post_type"`
	CategorySlug string `json:"category_slug" db:"category_slug"`
	Title        string `json:"title" db:"title"`
	Content      string `json:"content" db:"content"`
	Cover        string `json:"cover" db:"cover"`
	Download     string `json:"download" db:"download"`
	Type         string `json:"type" db:"type"`
	Video        string `json:"video" db:"video"`
}

type PostDetail struct {
	Owner interface{} `json:"owner"`
	*Post             // 帖子信息
}
