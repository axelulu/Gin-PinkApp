package mysql

import (
	"database/sql"
	"pinkacg/models"
	"strconv"
)

// CreateComment 创建评论
func CreateComment(uid int64, p *models.CommentCreate) (res sql.Result, err error) {
	// 创建评论
	sqlStr := `insert into comments (user_id, post_id, content, type, parent,update_time,create_time) values(?, ?, ?, ?, ?, NOW(), NOW())`

	pid, err := strconv.ParseInt(p.PostId, 10, 64)
	if err != nil {
		return nil, err
	}
	post, err := GetPostById(pid)
	if err != nil {
		return nil, err
	}
	// 修改文章评论数量
	sqlStr2 := `update posts set reply=?,update_time=NOW() where post_id=?`
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(sqlStr, uid, p.PostId, p.Content, p.Type, p.Parent)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}
	_, err = tx.Exec(sqlStr2, post.Reply+1, p.PostId)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	return
}

// GetCommentList 获取评论列表
func GetCommentList(p *models.CommentList) (comments []*models.Comment, err error) {
	sqlStr := `select user_id, post_id, content, type, parent, like_num, update_time from comments where post_id=? and status=1 limit ?,?`
	offset := (p.Page - 1) * p.Size
	err = db.Select(&comments, sqlStr, p.PostId, offset, p.Size)
	return
}
