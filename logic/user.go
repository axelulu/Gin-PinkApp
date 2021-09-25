package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2. 生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3. 保存进数据库

	if err := mysql.InsertUser(user); err != nil {
		return err
	}
	return
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 判断用户是否存在
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}

func UserById(id int64) (user *models.UserMeta, err error) {
	user, err = mysql.GetUserById(id)
	posts, err := mysql.GetPostsByUserId(id)
	user.Active = len(posts)
	return
}

func UserCenterById(uid int64) (userCenter map[string]interface{}, err error) {
	userCenter = make(map[string]interface{})
	user, err := mysql.GetUserMetaById(uid)
	userCenter["user"] = user
	// 文章
	posts, err := mysql.GetPostByPostType("post", uid, 2, 1)
	postCount, err := mysql.GetPostCountByPostType("post", uid)
	userCenter["posts"] = posts
	userCenter["postCount"] = postCount
	// 视频
	videos, err := mysql.GetPostByPostType("video", uid, 2, 1)
	videoCount, err := mysql.GetPostCountByPostType("video", uid)
	userCenter["videos"] = videos
	userCenter["videoCount"] = videoCount
	// 动态
	dynamics, err := mysql.GetPostByPostType("dynamic", uid, 2, 1)
	dynamicCount, err := mysql.GetPostCountByPostType("dynamic", uid)
	userCenter["dynamics"] = dynamics
	userCenter["dynamicCount"] = dynamicCount
	// 收藏
	var stars []*models.Star
	stars, err = mysql.GetStarsUserById(uid)
	if len(stars) > 0 {
		var starStr []int64
		for _, star := range stars {
			starStr = append(starStr, star.PostId)
		}
		var starPosts []*models.Post
		starPosts, err = mysql.GetPostByIds(starStr, 2, 1)
		starPostCount, _ := mysql.GetPostCountByIds(starStr)
		userCenter["starPosts"] = starPosts
		userCenter["starPostCount"] = starPostCount
	} else {
		userCenter["starPosts"] = nil
		userCenter["starPostCount"] = 0
	}
	// 喜欢
	var likes []*models.Like
	likes, err = mysql.GetLikesUserById(uid)
	if len(likes) > 0 {
		var likeStr []int64
		for _, like := range likes {
			likeStr = append(likeStr, like.PostId)
		}
		var likePosts []*models.Post
		likePosts, err = mysql.GetPostByIds(likeStr, 2, 1)
		likePostCount, _ := mysql.GetPostCountByIds(likeStr)
		userCenter["likePosts"] = likePosts
		userCenter["likePostCount"] = likePostCount
	} else {
		userCenter["likePosts"] = nil
		userCenter["likePostCount"] = 0
	}
	// 不喜欢
	var unLikes []*models.Like
	unLikes, err = mysql.GetUnLikesUserById(uid)
	if len(unLikes) > 0 {
		var unLikeStr []int64
		for _, unLike := range unLikes {
			unLikeStr = append(unLikeStr, unLike.PostId)
		}
		var unLikePosts []*models.Post
		unLikePosts, err = mysql.GetPostByIds(unLikeStr, 2, 1)
		unLikePostCount, _ := mysql.GetPostCountByIds(unLikeStr)
		userCenter["unLikePosts"] = unLikePosts
		userCenter["unLikePostCount"] = unLikePostCount
	} else {
		userCenter["unLikePosts"] = nil
		userCenter["unLikePostCount"] = 0
	}
	// 投币
	var coins []*models.Coin
	coins, err = mysql.GetCoinsUserById(uid)
	if len(coins) > 0 {
		var coinStr []int64
		for _, coin := range coins {
			coinStr = append(coinStr, coin.PostId)
		}
		var coinPosts []*models.Post
		coinPosts, err = mysql.GetPostByIds(coinStr, 2, 1)
		coinPostCount, _ := mysql.GetPostCountByIds(coinStr)
		userCenter["coinPosts"] = coinPosts
		userCenter["coinPostCount"] = coinPostCount
	} else {
		userCenter["coinPosts"] = nil
		userCenter["coinPostCount"] = 0
	}
	return
}
