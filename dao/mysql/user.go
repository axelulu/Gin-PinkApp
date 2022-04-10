package mysql

import (
	"database/sql"
	"pinkacg/models"

	"golang.org/x/crypto/bcrypt"
)

func CheckUserExist(email string) (count int, err error) {
	sqlStr := `select count(user_id) from users where email = ?`
	err = db.Get(&count, sqlStr, email)
	return
}

// InsertUser 向数据库中插入一条用户记录
func InsertUser(user *models.User) (err error) {
	password, err := hashBcrypt(user.Password)
	if err != nil {
		return err
	}

	sqlStr := `insert into users (user_id, username, email, password,is_vip,birth,update_time,create_time) values(?, ?, ?, ?, NOW(), NOW(),NOW(),NOW())`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Email, password)
	return
}

func hashBcrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", ErrorHashBcryptPassword
	}
	return string(hash), nil
}

func Login(user *models.User) (err error) {
	password := user.Password
	if err != nil {
		return err
	}
	sqlStr := `select user_id, email, password from users where email=?`
	if err = db.Get(user, sqlStr, user.Email); err == sql.ErrNoRows {
		return ErrorUserNotExist
	} else if err != nil {
		return err
	}

	//验证（对比）判读密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(uid int64) (user *models.UserMeta, err error) {
	user = new(models.UserMeta)
	sqlStr := `select user_id, username, email, fans, follows, coin, phone, avatar, background, descr, exp, gender, is_vip, birth from users where user_id=?`
	err = db.Get(user, sqlStr, uid)
	return
}

func GetUserMetaById(uid int64) (user *models.UserMeta, err error) {
	user = new(models.UserMeta)
	sqlStr := `select user_id, username, email, fans, follows, coin, phone, avatar, background, descr, exp, gender, is_vip, birth from users where user_id=?`
	err = db.Get(user, sqlStr, uid)
	return
}

func GetUserByUserMeta(word string, page int64, size int64) (user []*models.User, err error) {
	sqlStr := `select user_id, username, avatar, fans, follows, coin from users where username like '%` + word + `%' limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&user, sqlStr, offset, size)
	return
}

func UpdateUserInfo(uid int64, slugType string, slugVal string) (res sql.Result, err error) {
	sqlStr := "update users set " + slugType + "=?,update_time=NOW() where user_id = ?"
	res, err = db.Exec(sqlStr, slugVal, uid)
	return
}

func UpdateUserPassword(uid int64, p *models.UserPasswordUpdate) (res sql.Result, err error) {
	dbUser := new(models.User)
	sqlStr := `select user_id, email, password from users where email=?`
	if err = db.Get(dbUser, sqlStr, p.Email); err == sql.ErrNoRows {
		err = ErrorUserNotExist
		return
	}
	//验证（对比）判读密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(p.OldPassword)); err != nil {
		err = ErrorInvalidPassword
		return
	}
	newPassword, err := hashBcrypt(p.NewPassword)
	if err != nil {
		return nil, err
	}
	sqlStr2 := "update users set password=?,update_time=NOW() where user_id = ?"
	res, err = db.Exec(sqlStr2, newPassword, uid)
	return
}

func UpdateUserEmail(uid int64, p *models.UserEmailUpdate) (res sql.Result, err error) {
	sqlStr2 := "update users set email=?,update_time=NOW() where user_id = ?"
	res, err = db.Exec(sqlStr2, p.Email, uid)
	return
}

func UpdateUserPasswordByEmail(p *models.UserForgetPwd) (res sql.Result, err error) {
	sqlStr2 := "update users set password=?,update_time=NOW() where email = ?"
	newPassword, err := hashBcrypt(p.NewPassword)
	if err != nil {
		return nil, err
	}
	res, err = db.Exec(sqlStr2, newPassword, p.Email)
	return
}
