package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func GetPostDetail(posts []*models.Post) (postDetails []*models.PostDetail) {
	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("post.AuthorID", post.AuthorId), zap.Error(err))
			continue
		}
		postDetail := &models.PostDetail{
			Owner: user,
			Post:  post,
		}
		postDetails = append(postDetails, postDetail)
	}
	return
}

func HomeList(p *models.Home) (home map[string]interface{}, err error) {
	home = make(map[string]interface{}, 2)
	home["category"], err = mysql.GetCategoryList(p.CSize)
	posts, err := mysql.GetPostListByCategorySlug(p.CategorySlug, p.Size, p.Page)
	var postDetails []*models.PostDetail
	postDetails = GetPostDetail(posts)
	home["post"] = postDetails
	home["banner"] = postDetails
	return
}

func PostCategoryList(p *models.PostCategoryList) (post []*models.Post, err error) {
	post, err = mysql.GetPostListByCategorySlug(p.CategorySlug, p.Size, p.Page)
	return
}

func PostById(id int64) (postDetail map[string]interface{}, err error) {
	postDetail = make(map[string]interface{}, 4)
	var post *models.Post
	post, err = mysql.GetPostById(id)
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("post.AuthorID", post.AuthorId), zap.Error(err))
		return
	}
	postDetail["postInfo"] = &models.PostDetail{
		Owner: user,
		Post:  post,
	}
	posts, err := mysql.GetPostListByCategorySlug(post.CategorySlug, 10, 1)
	postDetail["postList"] = GetPostDetail(posts)
	postDetail["isFavorite"] = false
	postDetail["isLike"] = false
	return
}

func PostRanking(p *models.PostRankingList) (posts map[string]interface{}, err error) {
	posts = make(map[string]interface{}, 2)
	var post []*models.Post
	var postDetails []*models.PostDetail
	post, err = mysql.GetPostRanking(p.RankingSlug, p.Size, p.Page)
	postDetails = GetPostDetail(post)
	posts["list"] = postDetails
	posts["total"] = len(postDetails)
	return
}
