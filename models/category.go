package models

type CategoryList struct {
	Size int64 `json:"size" form:"size"`
}

type Category struct {
	CategorySlug string `json:"category_slug" db:"category_slug"`
	CategoryName string `json:"category_name" db:"category_name"`
}
