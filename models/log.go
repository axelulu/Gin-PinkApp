package models

type Log struct {
	Param    string `json:"param" form:"param"`
	ReadTime string `json:"readTime" form:"readTime"`
	//频道 分类ID
	CategoryId int64 `json:"categoryId" form:"categoryId"`
}

type LogParam struct {
	Action           string `json:"action"`
	UserId           string `json:"userId"`
	PostId           string `json:"postId"`
	AlgorithmCombine string `json:"algorithmCombine"`
}
