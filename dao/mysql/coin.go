package mysql

import (
	"database/sql"
	"web_app/models"
)

func GetCoinPostsById(pid int64, uid int64) (stars []*models.Star, err error) {
	sqlStr := `select user_id, post_id from coin where user_id=? and post_id=?`
	err = db.Select(&stars, sqlStr, uid, pid)
	return
}

func GetCoinsUserById(uid int64) (follows []*models.Coin, err error) {
	sqlStr := `select user_id, post_id from coin where user_id=?`
	err = db.Select(&follows, sqlStr, uid)
	return
}

func CoinPost(pid int64, uid int64, coin int64) (res sql.Result, err error) {
	// 获取用户硬币数量
	user := new(models.User)
	sqlStr2 := `select user_id, username, avatar, fans, follows, coin from user where user_id=?`
	err = db.Get(user, sqlStr2, uid)

	// 获取文章硬币数量
	post := new(models.Post)
	sqlStr5 := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where post_id=?`
	err = db.Get(post, sqlStr5, pid)

	// 开启事务
	tx, err := db.Begin()
	coins := user.Coin - coin
	if coin <= 0 {
		err = tx.Rollback()
	}
	sqlStr := `insert into coin (user_id, post_id, num) values(?,?,?)`
	res, err = tx.Exec(sqlStr, uid, pid, coin)
	if err != nil {
		err = tx.Rollback()
	}

	sqlStr3 := `update user set coin=? where user_id=?`
	res, err = tx.Exec(sqlStr3, coins, uid)
	if err != nil {
		err = tx.Rollback()
	}

	sqlStr4 := `update post set coin=? where post_id=?`
	res, err = tx.Exec(sqlStr4, post.Coin+coin, pid)
	if err != nil {
		err = tx.Rollback()
	}

	// 提交事务
	err = tx.Commit()
	return
}
