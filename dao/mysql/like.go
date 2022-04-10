package mysql

import (
	"database/sql"
	"pinkacg/models"
)

func GetLikePostsById(pid int64, uid int64, types int64) (likes []*models.Like, err error) {
	sqlStr := `select user_id, post_id, type from likes where user_id=? and post_id=? and type=?`
	err = db.Select(&likes, sqlStr, uid, pid, types)
	return
}

func GetLikesUserById(uid int64) (follows []*models.Like, err error) {
	sqlStr := `select user_id, post_id from likes where user_id=? and type=1`
	err = db.Select(&follows, sqlStr, uid)
	return
}

func GetUnLikesUserById(uid int64) (follows []*models.Like, err error) {
	sqlStr := `select user_id, post_id from likes where user_id=? and type=2`
	err = db.Select(&follows, sqlStr, uid)
	return
}

// UpdateLikePost 更新文章喜欢状态
func UpdateLikePost(pid int64, uid int64, types int64) (res sql.Result, err error) {
	// 获取喜欢的文章数量
	var likesNum []*models.Like
	sqlStr2 := `select user_id, post_id, type from likes where post_id=? and type=1`
	err = db.Select(&likesNum, sqlStr2, pid)
	if err != nil {
		return nil, err
	}

	// 获取不喜欢的文章数量
	var unLikesNum []*models.Like
	sqlStr3 := `select user_id, post_id, type from likes where user_id=? and post_id=? and type=2`
	err = db.Select(&unLikesNum, sqlStr3, uid, pid)
	if err != nil {
		return nil, err
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	sqlStr := `update likes set type=?,update_time=NOW() where user_id=? and post_id=?`
	res, err = tx.Exec(sqlStr, types, uid, pid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	sqlStr4 := `update posts set likes=?,un_likes=?,update_time=NOW() where post_id=?`

	if types == 1 && len(unLikesNum) > 0 {
		// 喜欢
		res, err = tx.Exec(sqlStr4, len(likesNum)+1, len(unLikesNum)-1, pid)
		if err != nil {
			if err = tx.Rollback(); err != nil {
				return nil, err
			}
		}
	} else if types == 2 && len(likesNum) > 0 {
		// 不喜欢
		res, err = tx.Exec(sqlStr4, len(likesNum)-1, len(unLikesNum)+1, pid)
		if err != nil {
			if err = tx.Rollback(); err != nil {
				return nil, err
			}
		}
	} else {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	// 提交事务
	err = tx.Commit()
	return
}

// LikePost 喜欢文章
// 0代表无状态，1代表喜欢，2代表不喜欢
func LikePost(pid int64, uid int64) (res sql.Result, err error) {
	// 获取喜欢的文章数量
	var likesNum []*models.Like
	sqlStr2 := `select user_id, post_id, type from likes where post_id=? and type=1`
	err = db.Select(&likesNum, sqlStr2, pid)
	if err != nil {
		return nil, err
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	sqlStr := `insert into likes (user_id, post_id, type,update_time,create_time) values(?,?,?,NOW(),NOW())`
	res, err = tx.Exec(sqlStr, uid, pid, 1)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}
	sqlStr3 := `update posts set likes=?,update_time=NOW() where post_id=?`
	res, err = tx.Exec(sqlStr3, len(likesNum)+1, pid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	// 提交事务
	err = tx.Commit()
	return
}

// UnLikePost 喜欢文章
// 0代表无状态，1代表喜欢，2代表不喜欢
func UnLikePost(pid int64, uid int64) (res sql.Result, err error) {
	// 获取喜欢的文章数量
	var unLikesNum []*models.Like
	sqlStr2 := `select user_id, post_id, type from likes where post_id=? and type=2`
	err = db.Select(&unLikesNum, sqlStr2, pid)
	if err != nil {
		return nil, err
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	sqlStr := `insert into likes (user_id, post_id, type,update_time,create_time) values(?,?,?,NOW(),NOW())`
	res, err = tx.Exec(sqlStr, uid, pid, 2)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	sqlStr3 := `update posts set un_likes=?,update_time=NOW() where post_id=?`
	res, err = tx.Exec(sqlStr3, len(unLikesNum)+1, pid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	// 提交事务
	err = tx.Commit()
	return
}
