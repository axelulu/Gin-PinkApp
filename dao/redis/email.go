package redis

import (
	"pinkacg/dao/mysql"
	"time"
)

const (
	KeyUserReg     = "user_reg_validate_code"
	KeyForgetPwd   = "forget_password_validate_code"
	KeyChangePwd   = "change_password_validate_code"
	KeyChangeEmail = "change_email_validate_code"
)

// SetEmailCode 设置验证码超时时间为5min
func SetEmailCode(slug, email, code string) (err error) {
	err = client.Set(GetRedisEmailKey(slug, email), code, 5*60*time.Second).Err()
	return
}

func CheckUserRegValidateCodeExist(slug, code string) (err error) {
	redisCode, err := client.Get(slug).Result()
	if redisCode != code {
		err = mysql.ErrorValidateCode
	}
	return
}
