package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func Search(p *models.Search) (content map[string]interface{}, err error) {
	content = make(map[string]interface{}, 2)
	if p.Type == "user" {
		var user []*models.User
		user, err = mysql.GetUserByUserMeta(p.Word, p.Page, p.Size)
		content["type"] = "user"
		content["user"] = user
	} else if p.Type == "post" || p.Type == "video" {
		var post []*models.Post
		post, err = mysql.GetPostByPostMeta(p.Word, p.Type, p.Page, p.Size)
		var postDetails []*models.PostDetail
		postDetails = GetPostDetail(post)
		content["type"] = "post"
		content["post"] = postDetails
	} else if p.Type == "all" {
		var post []*models.Post
		post, err = mysql.GetAllPostByPostMeta(p.Word, p.Page, p.Size)
		var postDetails []*models.PostDetail
		postDetails = GetPostDetail(post)
		content["type"] = "post"
		content["post"] = postDetails
	}
	return
}
