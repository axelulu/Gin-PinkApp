package mysql

import (
	"database/sql"
	"pinkacg/models"
)

func GetCoinPostsById(pid int64, uid int64) (stars []*models.Star, err error) {
	sqlStr := `select user_id, post_id from coins where user_id=? and post_id=?`
	err = db.Select(&stars, sqlStr, uid, pid)
	return
}

func GetCoinsUserById(uid int64) (follows []*models.Coin, err error) {
	sqlStr := `select user_id, post_id from coins where user_id=?`
	err = db.Select(&follows, sqlStr, uid)
	return
}

func CoinPost(pid int64, uid int64, coin int64) (res sql.Result, err error) {
	// 获取用户硬币数量
	user := new(models.User)
	sqlStr2 := `select user_id, username, avatar, fans, follows, coin from users where user_id=?`
	err = db.Get(user, sqlStr2, uid)
	if err != nil {
		return nil, err
	}

	// 获取文章硬币数量
	post := new(models.Post)
	sqlStr5 := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where post_id=?`
	err = db.Get(post, sqlStr5, pid)
	if err != nil {
		return nil, err
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	coins := user.Coin - coin
	if coin <= 0 || coins < 0 {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}
	sqlStr := `insert into coins (user_id, post_id, num, update_time, create_time) values(?,?,?,NOW(),NOW())`
	res, err = tx.Exec(sqlStr, uid, pid, coin)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	sqlStr3 := `update users set coin=?,update_time=NOW() where user_id=?`
	res, err = tx.Exec(sqlStr3, coins, uid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	sqlStr4 := `update posts set coin=?,update_time=NOW() where post_id=?`
	res, err = tx.Exec(sqlStr4, post.Coin+coin, pid)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return nil, err
		}
	}

	// 提交事务
	err = tx.Commit()
	return
}
