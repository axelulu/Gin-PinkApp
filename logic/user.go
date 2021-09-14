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

func UserById(id int64) (user *models.User, err error) {
	user, err = mysql.GetUserById(id)
	return
}
