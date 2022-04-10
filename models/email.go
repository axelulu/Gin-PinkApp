package models

type Email struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}
