package logic

import (
	"database/sql"
	"pinkacg/dao/mysql"
	"pinkacg/dao/redis"
	"pinkacg/models"
	"pinkacg/pkg/jwt"
	"pinkacg/pkg/snowflake"
)

// SignUp 注册
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if count, _ := mysql.CheckUserExist(p.Email); count > 0 {
		err = mysql.ErrorEmailExist
		return
	}
	// 判断验证码是否存在
	if err = redis.CheckUserRegValidateCodeExist(redis.GetRedisEmailKey(redis.KeyUserReg, p.Email), p.ValidateCode); err != nil {
		return mysql.ErrorValidateCode
	}
	// 生成UID
	userID := snowflake.GenID()
	// 构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	}
	// 保存进数据库
	err = mysql.InsertUser(user)
	return
}

// Login 登录
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Email:    p.Email,
		Password: p.Password,
	}
	// 判断用户是否存在
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}

func ForgetPwd(p *models.UserForgetPwd) (res sql.Result, err error) {
	// 1. 判断用户是否存在
	count, err := mysql.CheckUserExist(p.Email)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, mysql.ErrorUserNotExist
	}
	// 判断验证码是否存在
	if err = redis.CheckUserRegValidateCodeExist(redis.GetRedisEmailKey(redis.KeyForgetPwd, p.Email), p.ValidateCode); err != nil {
		return nil, mysql.ErrorValidateCode
	}
	res, err = mysql.UpdateUserPasswordByEmail(p)
	return
}

// UserById 根据ID获取用户信息
func UserById(id int64) (user *models.UserMeta, err error) {
	user, err = mysql.GetUserById(id)
	if err != nil {
		return nil, mysql.ErrorUserMeta
	}
	posts, err := mysql.GetPostCountByPostType("dynamic", id)
	if err != nil {
		return nil, err
	}
	user.Active = posts
	return
}

// UserCenterById 根据用户ID获取用户中心
func UserCenterById(uid int64) (userCenter map[string]interface{}, err error) {
	userCenter = make(map[string]interface{})
	user, err := mysql.GetUserById(uid)
	if err != nil {
		return nil, err
	}
	dynamicPosts, err := mysql.GetPostByPostTypeAndUserID("dynamic", uid, 1, 10000)
	if err != nil {
		return nil, err
	}
	user.Active = int64(len(dynamicPosts))
	userCenter["user"] = user
	// 文章
	posts, err := mysql.GetPostByPostType("post", uid, 2, 1)
	if err != nil {
		return nil, err
	}
	postCount, err := mysql.GetPostCountByPostType("post", uid)
	if err != nil {
		return nil, err
	}
	var postDetails []*models.PostDetail
	postDetails = GetPostDetail(posts)
	userCenter["posts"] = postDetails
	userCenter["postCount"] = postCount
	// 视频
	videos, err := mysql.GetPostByPostType("video", uid, 2, 1)
	if err != nil {
		return nil, err
	}
	videoCount, err := mysql.GetPostCountByPostType("video", uid)
	if err != nil {
		return nil, err
	}
	var videoDetails []*models.PostDetail
	videoDetails = GetPostDetail(videos)
	userCenter["videos"] = videoDetails
	userCenter["videoCount"] = videoCount
	// 动态
	dynamics, err := mysql.GetPostByPostType("dynamic", uid, 2, 1)
	if err != nil {
		return nil, err
	}
	dynamicCount, err := mysql.GetPostCountByPostType("dynamic", uid)
	if err != nil {
		return nil, err
	}
	var dynamicDetails []*models.PostDetail
	dynamicDetails = GetPostDetail(dynamics)
	userCenter["dynamics"] = dynamicDetails
	userCenter["dynamicCount"] = dynamicCount
	// 收藏
	userCenter["starPosts"], userCenter["starPostCount"], err = GetStarPost(uid, 2, 1)
	if err != nil {
		return nil, err
	}
	// 喜欢
	userCenter["likePosts"], userCenter["likePostCount"], err = GetLikePost(uid, 2, 1)
	if err != nil {
		return nil, err
	}
	// 不喜欢
	userCenter["unLikePosts"], userCenter["unLikePostCount"], err = GetUnLikePost(uid, 2, 1)
	if err != nil {
		return nil, err
	}
	// 投币
	userCenter["coinPosts"], userCenter["coinPostCount"], err = GetCoinPost(uid, 2, 1)
	if err != nil {
		return nil, err
	}
	return
}

// UserInfoUpdate 用户信息更新
func UserInfoUpdate(uid int64, p *models.UserUpdate) (res sql.Result, err error) {
	res, err = mysql.UpdateUserInfo(uid, p.Slug, p.Value)
	return
}

// UserPasswordUpdate 用户密码更新
func UserPasswordUpdate(uid int64, p *models.UserPasswordUpdate) (res sql.Result, err error) {
	if err = redis.CheckUserRegValidateCodeExist(redis.GetRedisEmailKey(redis.KeyChangePwd, p.Email), p.ValidateCode); err != nil {
		return nil, mysql.ErrorValidateCode
	}
	// 1. 判断邮箱是否存在
	res, err = mysql.UpdateUserPassword(uid, p)
	return
}

// UserEmailUpdate 用户邮箱更新
func UserEmailUpdate(uid int64, p *models.UserEmailUpdate) (res sql.Result, err error) {
	// 1. 判断邮箱是否存在
	count, err := mysql.CheckUserExist(p.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = mysql.ErrorEmailExist
		return
	}
	if err = redis.CheckUserRegValidateCodeExist(redis.GetRedisEmailKey(redis.KeyChangeEmail, p.Email), p.ValidateCode); err != nil {
		return nil, mysql.ErrorValidateCode
	}
	// 1. 判断邮箱是否存在
	res, err = mysql.UpdateUserEmail(uid, p)
	return
}
