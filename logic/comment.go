package logic

import (
	"database/sql"
	"go.uber.org/zap"
	"pinkacg/dao/mysql"
	"pinkacg/models"
)

// CommentCreate 评论创建
func CommentCreate(uid int64, p *models.CommentCreate) (res sql.Result, err error) {
	res, err = mysql.CreateComment(uid, p)
	return
}

// CommentList 获取评论列表
func CommentList(p *models.CommentList) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 2)
	var comments []*models.Comment
	var commentDetails []*models.CommentDetail
	comments, err = mysql.GetCommentList(p)
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(comment.UserId)
		if err != nil {
			zap.L().Error("mysql.GetUserById failed", zap.Int64("comment.UserId", comment.UserId), zap.Error(err))
			return nil, mysql.ErrorUserMeta
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
