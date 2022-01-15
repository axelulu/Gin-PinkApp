package logic

import (
	"pinkacg/dao/mysql"
	"pinkacg/dao/redis"
	"pinkacg/models"
	"pinkacg/pkg/email"
)

// UserRegSendEmail 发送用户注册邮件
func UserRegSendEmail(p *models.Email, code string) (err error) {
	// 1. 判断用户是否存在
	count, err := mysql.CheckUserExist(p.Email)
	if err != nil {
		return err
	}
	if count > 0 {
		return mysql.ErrorEmailExist
	}
	if err = email.SendEmail("reg new user", p.Email, "code:"+code); err != nil {
		return
	}
	err = redis.SetEmailCode(redis.KeyUserReg, p.Email, code)
	return
}

// UserForgetPwdSendEmail 发送忘记密码邮件
func UserForgetPwdSendEmail(p *models.Email, code string) (err error) {
	// 1. 判断用户是否存在
	count, err := mysql.CheckUserExist(p.Email)
	if err != nil {
		return err
	}
	if count == 0 {
		return mysql.ErrorUserNotExist
	}
	if err = email.SendEmail("forget password", p.Email, "code:"+code); err != nil {
		return
	}
	err = redis.SetEmailCode(redis.KeyForgetPwd, p.Email, code)
	return
}

// UserChangePwdSendEmail 发送改变密码邮件
func UserChangePwdSendEmail(p *models.Email, code string, uid int64) (err error) {
	// 1. 判断用户是否存在
	user, _ := mysql.GetUserById(uid)
	if user.Email != p.Email {
		return mysql.ErrorUserNotExist
	}
	if err = email.SendEmail("change password", p.Email, "code:"+code); err != nil {
		return
	}
	err = redis.SetEmailCode(redis.KeyChangePwd, p.Email, code)
	return
}

// UserChangeEmailSendEmail 发送改变邮箱邮件
func UserChangeEmailSendEmail(p *models.Email, code string) (err error) {
	// 1. 判断邮箱是否存在
	count, err := mysql.CheckUserExist(p.Email)
	if err != nil {
		return err
	}
	if count > 0 {
		return mysql.ErrorEmailExist
	}
	if err = email.SendEmail("change email", p.Email, "code:"+code); err != nil {
		return
	}
	err = redis.SetEmailCode(redis.KeyChangeEmail, p.Email, code)
	return
}
