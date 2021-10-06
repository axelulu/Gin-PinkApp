package logic

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// GetPostDetail 获取文章详细信息
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

// HomeList 首页文章列表
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

// PostCategoryList 根据分类获取文章列表
func PostCategoryList(p *models.PostCategoryList) (post []*models.Post, err error) {
	post, err = mysql.GetPostListByCategorySlug(p.CategorySlug, p.Size, p.Page)
	return
}

// PostById 根据文章id获取文章
func PostById(id int64, uid int64) (postDetail map[string]interface{}, err error) {
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
	follows, err := mysql.GetFollowUserById(post.AuthorId, uid)
	if len(follows) > 0 {
		postDetail["isFollow"] = true
	} else {
		postDetail["isFollow"] = false
	}
	postDetail["isSelf"] = post.AuthorId == uid
	// 是否收藏文章
	star, err := mysql.GetStarPostsById(id, uid)
	if len(star) > 0 {
		postDetail["isFavorite"] = true
	} else {
		postDetail["isFavorite"] = false
	}
	// 是否投币文章
	coin, err := mysql.GetCoinPostsById(id, uid)
	if len(coin) > 0 {
		postDetail["isCoin"] = true
	} else {
		postDetail["isCoin"] = false
	}
	// 不喜欢文章
	like, err := mysql.GetLikePostsById(id, uid, 2)
	if len(like) > 0 {
		postDetail["isUnLike"] = true
	} else {
		postDetail["isUnLike"] = false
	}
	// 已经喜欢文章
	likes, err := mysql.GetLikePostsById(id, uid, 1)
	if len(likes) > 0 {
		postDetail["isLike"] = true
	} else {
		postDetail["isLike"] = false
	}
	return
}

// PostPublish 发布文章
func PostPublish(p *models.PostPublish, authorId int64) (exec sql.Result, err error) {
	postId := snowflake.GenID()
	exec, err = mysql.CreatePost(postId, authorId, p.PostType, p.CategorySlug, p.Title, p.Cover, p.Content, p.Video)
	return
}

// PostRanking 文章排行
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

// PostDynamic 文章动态
func PostDynamic(p *models.PostDynamicList, uid int64) (posts map[string]interface{}, err error) {
	var follows []*models.Follow
	follows, err = mysql.GetFollowsUserById(uid)
	posts = make(map[string]interface{}, 2)
	var post []*models.Post
	var postDetails []*models.PostDetail
	var uids []int64
	for _, user := range follows {
		uids = append(uids, user.FollowId)
	}
	uids = append(uids, uid)
	post, err = mysql.GetPostDynamicByIds(uids, p.Size, p.Page, p.DynamicSlug)
	postDetails = GetPostDetail(post)
	posts["list"] = postDetails
	posts["total"] = len(postDetails)
	return
}

// LikePost 喜欢文章
func LikePost(pid int64, uid int64) (res sql.Result, err error) {
	// 不喜欢文章
	like, err := mysql.GetLikePostsById(pid, uid, 2)
	// 已经喜欢文章
	likes, err := mysql.GetLikePostsById(pid, uid, 1)
	if err != nil {
		return nil, err
	}
	if len(like) > 0 {
		// 该文章被不喜欢，改为喜欢
		res, err = mysql.UpdateLikePost(pid, uid, 1)
	} else if len(likes) > 0 {
		// 该文章已经被喜欢
		err = errors.New("已经喜欢了该文章")
	} else {
		// 该文章无状态，喜欢该文章
		res, err = mysql.LikePost(pid, uid)
	}
	return
}

// UnLikePost 不喜欢文章
func UnLikePost(pid int64, uid int64) (res sql.Result, err error) {
	// 不喜欢文章
	like, err := mysql.GetLikePostsById(pid, uid, 2)
	// 已经喜欢文章
	likes, err := mysql.GetLikePostsById(pid, uid, 1)
	if err != nil {
		return nil, err
	}
	if len(like) > 0 {
		// 该文章已经被喜欢
		err = errors.New("已经不喜欢了该文章")
	} else if len(likes) > 0 {
		// 该文章被喜欢，改为不喜欢
		res, err = mysql.UpdateLikePost(pid, uid, 2)
	} else {
		// 该文章无状态，不喜欢该文章
		res, err = mysql.UnLikePost(pid, uid)
	}
	return
}

// StarPost 收藏文章
func StarPost(pid int64, uid int64) (res sql.Result, err error) {
	// 是否收藏文章
	star, err := mysql.GetStarPostsById(pid, uid)
	if err != nil {
		return nil, err
	}

	if len(star) > 0 {
		// 该文章已经被收藏
		err = errors.New("已经收藏了该文章")
	} else {
		// 该文章无状态，收藏该文章
		res, err = mysql.StarPost(pid, uid)
	}
	return
}

// UnStarPost 不收藏文章
func UnStarPost(pid int64, uid int64) (res sql.Result, err error) {
	// 是否收藏文章
	star, err := mysql.GetStarPostsById(pid, uid)
	if err != nil {
		return nil, err
	}

	if len(star) > 0 {
		// 该文章被收藏，不收藏该文章
		res, err = mysql.UnStarPost(pid, uid)
	} else {
		// 该文章已经被收藏
		err = errors.New("没有收藏该文章")
	}
	return
}

// CoinPost 不收藏文章
func CoinPost(pid int64, uid int64, coin int64) (res sql.Result, err error) {
	// 是否收藏文章
	coins, err := mysql.GetCoinPostsById(pid, uid)
	if err != nil {
		return nil, err
	}

	if len(coins) > 0 {
		// 该文章已投币
		err = errors.New("该文章已投币")
	} else {
		// 该文章未被投币
		res, err = mysql.CoinPost(pid, uid, coin)
	}
	return
}

// GetUserPost 根据用户id与文章类型获取文章
func GetUserPost(p *models.UserPost, uid int64) (posts map[string]interface{}, err error) {
	var post []*models.Post
	var postDetails []*models.PostDetail
	posts = make(map[string]interface{}, 2)
	if p.PostType == "star" {
		posts["list"], posts["total"], err = GetStarPost(uid, p.Size, p.Page)
	} else if p.PostType == "coin" {
		posts["list"], posts["total"], err = GetCoinPost(uid, p.Size, p.Page)
	} else if p.PostType == "like" {
		posts["list"], posts["total"], err = GetLikePost(uid, p.Size, p.Page)
	} else if p.PostType == "unlike" {
		posts["list"], posts["total"], err = GetUnLikePost(uid, p.Size, p.Page)
	} else {
		post, err = mysql.GetPostByPostTypeAndUserID(p.PostType, p.UserId, p.Page, p.Size)
		postDetails = GetPostDetail(post)
		posts["list"] = postDetails
		posts["total"] = len(postDetails)
	}
	return
}

func GetStarPost(uid int64, count int64, page int64) (starPostsDetails []*models.PostDetail, starPostCount int64, err error) {
	var stars []*models.Star
	stars, err = mysql.GetStarUserById(uid)
	if len(stars) > 0 {
		var starStr []int64
		for _, star := range stars {
			starStr = append(starStr, star.PostId)
		}
		var starPosts []*models.Post
		starPosts, err = mysql.GetPostByIds(starStr, count, page)
		starPostCount, _ = mysql.GetPostCountByIds(starStr)
		starPostsDetails = GetPostDetail(starPosts)
	} else {
		starPostsDetails = nil
		starPostCount = 0
	}
	return
}

func GetLikePost(uid int64, count int64, page int64) (likePostsDetails []*models.PostDetail, likePostCount int64, err error) {
	var likes []*models.Like
	likes, err = mysql.GetLikesUserById(uid)
	if len(likes) > 0 {
		var likeStr []int64
		for _, like := range likes {
			likeStr = append(likeStr, like.PostId)
		}
		var likePosts []*models.Post
		likePosts, err = mysql.GetPostByIds(likeStr, count, page)
		likePostCount, _ = mysql.GetPostCountByIds(likeStr)
		likePostsDetails = GetPostDetail(likePosts)
	}
	return
}

func GetUnLikePost(uid int64, count int64, page int64) (unLikePostsDetails []*models.PostDetail, unLikePostCount int64, err error) {
	var unLikes []*models.Like
	unLikes, err = mysql.GetUnLikesUserById(uid)
	if len(unLikes) > 0 {
		var unLikeStr []int64
		for _, unLike := range unLikes {
			unLikeStr = append(unLikeStr, unLike.PostId)
		}
		var unLikePosts []*models.Post
		unLikePosts, err = mysql.GetPostByIds(unLikeStr, count, page)
		unLikePostCount, _ = mysql.GetPostCountByIds(unLikeStr)
		unLikePostsDetails = GetPostDetail(unLikePosts)
	}
	return
}

func GetCoinPost(uid int64, count int64, page int64) (coinPostsDetails []*models.PostDetail, coinPostCount int64, err error) {
	var coins []*models.Coin
	coins, err = mysql.GetCoinsUserById(uid)
	if len(coins) > 0 {
		var coinStr []int64
		for _, coin := range coins {
			coinStr = append(coinStr, coin.PostId)
		}
		var coinPosts []*models.Post
		coinPosts, err = mysql.GetPostByIds(coinStr, count, page)
		coinPostCount, _ = mysql.GetPostCountByIds(coinStr)
		coinPostsDetails = GetPostDetail(coinPosts)
	}
	return
}
