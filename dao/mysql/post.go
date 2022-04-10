package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"pinkacg/models"
)

func GetRecommendPostList(count int64, page int64, sort string) (post []*models.Post, err error) {
	var sqlStr string
	if sort == "rand" {
		sqlStr = "select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts ORDER BY RAND() desc limit ?,?"
	} else {
		sqlStr = "select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts ORDER BY " + sort + " desc limit ?,?"
	}
	start := (page - 1) * count
	err = db.Select(&post, sqlStr, start, count)
	return
}

func GetPostListByCategorySlug(categorySlug int64, count int64, page int64, sort string) (post []*models.Post, err error) {
	var sqlStr string
	if sort == "rand" {
		sqlStr = "select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where category_id=? ORDER BY RAND() desc limit ?,?"
	} else {
		sqlStr = "select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where category_id=? ORDER BY " + sort + " desc limit ?,?"
	}
	start := (page - 1) * count
	err = db.Select(&post, sqlStr, categorySlug, start, count)
	return
}

func GetPostById(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where post_id=?`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostByCId(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where id=?`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostRanking(rankingSlug string, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := "select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts order by " + rankingSlug + " DESC limit ?,?"
	offset := (page - 1) * count
	err = db.Select(&post, sqlStr, offset, count)
	return
}

func GetPostByPostType(postType string, userId int64, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where post_type=? and author_id=? ORDER BY update_time desc limit ?,?`
	offset := (page - 1) * count
	err = db.Select(&post, sqlStr, postType, userId, offset, count)
	return
}

func GetPostCountByPostType(postType string, userId int64) (count int64, err error) {
	sqlStr := `select count(*) from posts where post_type=? and author_id=?`
	err = db.QueryRow(sqlStr, postType, userId).Scan(&count)
	return
}

func GetPostDynamicByIds(uids []int64, count int64, page int64, dynamicSlug string) (post []*models.Post, err error) {
	offset := (page - 1) * count
	var query string
	var args []interface{}
	if dynamicSlug == "all" {
		sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where author_id in(?) ORDER BY update_time desc limit ?,?`
		query, args, err = sqlx.In(sqlStr, uids, offset, count)
		if err != nil {
			return nil, err
		}
	} else {
		sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where post_type=? and author_id in(?) ORDER BY update_time desc limit ?,?`
		query, args, err = sqlx.In(sqlStr, dynamicSlug, uids, offset, count)
		if err != nil {
			return nil, err
		}
	}
	query = db.Rebind(query) // Rebind query
	err = db.Select(&post, query, args...)
	return
}

func GetPostByIds(uids interface{}, count int64, page int64) (post []*models.Post, err error) {
	offset := (page - 1) * count
	var query string
	var args []interface{}
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where post_id in(?) ORDER BY null limit ?,?`
	query, args, err = sqlx.In(sqlStr, uids, offset, count)
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query) // Rebind query
	err = db.Select(&post, query, args...)
	return
}

func GetPostCountByIds(uids []int64) (count int64, err error) {
	sqlStr := `select count(*) from posts where post_id in(?)`
	query, args, err := sqlx.In(sqlStr, uids)
	if err != nil {
		return 0, err
	}
	query = db.Rebind(query) // Rebind query
	err = db.QueryRow(query, args...).Scan(&count)
	return
}

func GetPostByPostMeta(word string, postType string, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where post_type=? and (title like '%` + word + `%' or content like '%` + word + `%') ORDER BY update_time desc limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, postType, offset, size)
	return
}

func GetAllPostByPostMeta(word string, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where (title like '%` + word + `%' or content like '%` + word + `%') ORDER BY update_time desc limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, offset, size)
	return
}

func GetPostByPostTypeAndUserID(postType string, userId int64, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_id, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from posts where author_id=? and post_type=? ORDER BY update_time desc limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, userId, postType, offset, size)
	return
}

func CreatePost(postId int64, authorId int64, postType string, categorySlug int64, title string, cover string, content string, video string) (exec sql.Result, err error) {
	sqlStr := `insert into posts (post_id,author_id,post_type,category_id,title,cover,content,video,update_time,create_time) values(?,?,?,?,?,?,?,?,NOW(),NOW())`
	exec, err = db.Exec(sqlStr, postId, authorId, postType, categorySlug, title, cover, content, video)
	return
}

func AddPostViewByPostId(postId int64) (err error) {
	sqlStr := `update posts set view=? where post_id=?`
	post, err := GetPostById(postId)
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlStr, post.View+1, postId)
	return
}

func UpdatePostByPostId(postId int64, slug string, value string) (err error) {
	sqlStr := "update posts set " + slug + "=? where post_id=?"
	_, err = GetPostById(postId)
	if err != nil {
		return err
	}
	_, err = db.Exec(sqlStr, value, postId)
	return
}
