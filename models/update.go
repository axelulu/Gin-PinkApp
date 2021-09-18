package models

type Update struct {
	Version string `json:"version" form:"version"`
	Url     string `json:"url" form:"url"`
	Meta    string `json:"meta" form:"meta"`
}
