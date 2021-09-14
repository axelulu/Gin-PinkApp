package mysql

import (
	"database/sql"
	"web_app/models"

	"golang.org/x/crypto/bcrypt"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条用户记录
func InsertUser(user *models.User) (err error) {
	password, err := hashBcrypt(user.Password)
	if err != nil {
		return err
	}

	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	if err != nil {
		return err
	}
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
	sqlStr := `select user_id, username, password from user where username=?`
	if err = db.Get(user, sqlStr, user.Username); err == sql.ErrNoRows {
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

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username, avatar, fans from user where user_id=?`
	err = db.Get(user, sqlStr, uid)
	return
}

func GetUserByUserMeta(word string, page int64, size int64) (user []*models.User, err error) {
	sqlStr := `select user_id, username, avatar, fans from user where username like '%` + word + `%' limit ?,?`
	offset := (page - 1) * size
	err = db.Select(&user, sqlStr, offset, size)
	return
}
