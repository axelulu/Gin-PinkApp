package mysql

import (
	"database/sql"
	"web_app/models"
)

func CreateComment(uid int64, p *models.CommentCreate) (res sql.Result, err error) {
	sqlStr := `insert into comment (user_id, post_id, content, type, parent) values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, uid, p.PostId, p.Content, p.Type, p.Parent)
	return
}

func GetCommentList(p *models.CommentList) (comments []*models.Comment, err error) {
	sqlStr := `select user_id, post_id, content, type, parent, like_num, updated_time from comment where post_id=? and status=1 limit ?,?`
	offset := (p.Page - 1) * p.Size
	err = db.Select(&comments, sqlStr, p.PostId, offset, p.Size)
	return
}
