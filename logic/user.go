package logic

import (
	"database/sql"
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
	var postDetails []*models.PostDetail
	postDetails = GetPostDetail(posts)
	userCenter["posts"] = postDetails
	userCenter["postCount"] = postCount
	// 视频
	videos, err := mysql.GetPostByPostType("video", uid, 2, 1)
	videoCount, err := mysql.GetPostCountByPostType("video", uid)
	var videoDetails []*models.PostDetail
	videoDetails = GetPostDetail(videos)
	userCenter["videos"] = videoDetails
	userCenter["videoCount"] = videoCount
	// 动态
	dynamics, err := mysql.GetPostByPostType("dynamic", uid, 2, 1)
	dynamicCount, err := mysql.GetPostCountByPostType("dynamic", uid)
	var dynamicDetails []*models.PostDetail
	dynamicDetails = GetPostDetail(dynamics)
	userCenter["dynamics"] = dynamicDetails
	userCenter["dynamicCount"] = dynamicCount
	// 收藏
	userCenter["starPosts"], userCenter["starPostCount"], err = GetStarPost(uid, 2, 1)
	// 喜欢
	userCenter["likePosts"], userCenter["likePostCount"], err = GetLikePost(uid, 2, 1)
	// 不喜欢
	userCenter["unLikePosts"], userCenter["unLikePostCount"], err = GetUnLikePost(uid, 2, 1)
	// 投币
	userCenter["coinPosts"], userCenter["coinPostCount"], err = GetCoinPost(uid, 2, 1)
	return
}

func UserInfoUpdate(uid int64, p *models.UserUpdate) (res sql.Result, err error) {
	res, err = mysql.UpdateUserInfo(uid, p.Slug, p.Value)
	return
}
