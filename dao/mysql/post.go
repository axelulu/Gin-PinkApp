package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"web_app/models"
)

func GetPostListByCategorySlug(categorySlug string, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where category_slug=? limit ?,?`
	start := (page - 1) * count
	err = db.Select(&post, sqlStr, categorySlug, start, count)
	return
}

func GetPostById(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where post_id=?`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostsByUserId(uid int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where author_id=?`
	err = db.Select(&posts, sqlStr, uid)
	return
}

func GetPostRanking(rankingSlug string, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post order by ? limit ?,?`
	offset := (page - 1) * count
	err = db.Select(&post, sqlStr, rankingSlug, offset, count)
	print(len(post))
	return
}

func GetPostByPostType(postType string, userId int64, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where post_type=? and author_id=? limit ?,?`
	offset := (page - 1) * count
	err = db.Select(&post, sqlStr, postType, userId, offset, count)
	return
}

func GetPostCountByPostType(postType string, userId int64) (count int64, err error) {
	sqlStr := `select count(*) from post where post_type=? and author_id=?`
	err = db.QueryRow(sqlStr, postType, userId).Scan(&count)
	return
}

func GetPostDynamicByIds(uids []int64, count int64, page int64, dynamicSlug string) (post []*models.Post, err error) {
	offset := (page - 1) * count
	var query string
	var args []interface{}
	if dynamicSlug == "all" {
		sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where author_id in(?) limit ?,?`
		query, args, err = sqlx.In(sqlStr, uids, offset, count)
	} else {
		sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where post_type=? and author_id in(?) limit ?,?`
		query, args, err = sqlx.In(sqlStr, dynamicSlug, uids, offset, count)
	}
	query = db.Rebind(query) // Rebind query
	err = db.Select(&post, query, args...)
	return
}

func GetPostByIds(uids []int64, count int64, page int64) (post []*models.Post, err error) {
	offset := (page - 1) * count
	var query string
	var args []interface{}
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where post_id in(?) limit ?,?`
	query, args, err = sqlx.In(sqlStr, uids, offset, count)
	query = db.Rebind(query) // Rebind query
	err = db.Select(&post, query, args...)
	return
}

func GetPostCountByIds(uids []int64) (count int64, err error) {
	sqlStr := `select count(*) from post where post_id in(?)`
	query, args, err := sqlx.In(sqlStr, uids)
	query = db.Rebind(query) // Rebind query
	err = db.QueryRow(query, args...).Scan(&count)
	return
}

func GetPostByPostMeta(word string, postType string, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where post_type=? and (title like '%` + word + `%' or content like '%` + word + `%') limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, postType, offset, size)
	return
}

func GetAllPostByPostMeta(word string, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where (title like '%` + word + `%' or content like '%` + word + `%') limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, offset, size)
	return
}

func GetPostByPostTypeAndUserID(postType string, userId int64, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, un_likes, coin, share, view, cover, video, download, create_time, update_time from post where author_id=? and post_type=? limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, userId, postType, offset, size)
	return
}

func CreatePost(postId int64, authorId int64, postType string, categorySlug string, title string, cover string, content string, video string) (exec sql.Result, err error) {
	sqlStr := `insert into post (post_id,author_id,post_type,category_slug,title,cover,content,video) values(?,?,?,?,?,?,?,?)`
	exec, err = db.Exec(sqlStr, postId, authorId, postType, categorySlug, title, cover, content, video)
	return
}
