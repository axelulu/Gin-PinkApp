package mysql

import (
	"database/sql"
	"web_app/models"
)

func GetPostListByCategorySlug(categorySlug string, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, coin, share, view, cover, video, download, create_time, update_time from post where category_slug=? limit ?,?`
	start := (page - 1) * count
	err = db.Select(&post, sqlStr, categorySlug, start, count)
	return
}

func GetPostById(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, coin, share, view, cover, video, download, create_time, update_time from post where post_id=?`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostRanking(rankingSlug string, count int64, page int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, coin, share, view, cover, video, download, create_time, update_time from post order by ? limit ?,?`
	offset := (page - 1) * count
	err = db.Select(&post, sqlStr, rankingSlug, offset, count)
	return
}

func GetPostByPostMeta(word string, postType string, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, coin, share, view, cover, video, download, create_time, update_time from post where post_type=? and (title like '%` + word + `%' or content like '%` + word + `%') limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, postType, offset, size)
	return
}

func GetAllPostByPostMeta(word string, page int64, size int64) (post []*models.Post, err error) {
	sqlStr := `select post_id, author_id, post_type, category_slug, title, content, reply, favorite, likes, coin, share, view, cover, video, download, create_time, update_time from post where (title like '%` + word + `%' or content like '%` + word + `%') limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&post, sqlStr, offset, size)
	return
}

func CreatePost(postId int64, authorId int64, postType string, categorySlug string, title string, cover string, content string, video string) (exec sql.Result, err error) {
	sqlStr := `insert into post (post_id,author_id,post_type,category_slug,title,cover,content,video) values(?,?,?,?,?,?,?,?)`
	exec, err = db.Exec(sqlStr, postId, authorId, postType, categorySlug, title, cover, content, video)
	return
}
