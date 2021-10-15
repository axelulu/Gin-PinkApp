package logic

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func CommentCreate(uid int64, p *models.CommentCreate) (res sql.Result, err error) {
	res, err = mysql.CreateComment(uid, p)
	return
}

func CommentList(p *models.CommentList) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 2)
	var comments []*models.Comment
	var commentDetails []*models.CommentDetail
	comments, err = mysql.GetCommentList(p)
	for _, comment := range comments {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(comment.UserId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("post.AuthorID", comment.UserId), zap.Error(err))
			continue
		}
		postDetail := &models.CommentDetail{
			Owner:   user,
			Comment: comment,
		}
		commentDetails = append(commentDetails, postDetail)
	}
	result["list"] = commentDetails
	result["total"] = len(commentDetails)
	return
}
