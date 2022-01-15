package models

type CategoryList struct {
	Size int64 `json:"size" form:"size"`
}

type Category struct {
	CategorySlug int64  `json:"category_id" db:"category_id"`
	CategoryName string `json:"category_name" db:"category_name"`
}
