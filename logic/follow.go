package logic

import (
	"database/sql"
	"errors"
	"web_app/dao/mysql"
)

func FollowUser(followId int64, uid int64) (res sql.Result, err error) {
	id, err := mysql.GetFollowUserById(followId, uid)
	if err != nil {
		return nil, err
	}
	if len(id) == 0 {
		res, err = mysql.InsertFollowUserById(followId, uid)
	} else {
		err = errors.New("已经关注此用户！")
	}
	return
}

func UnFollowUser(followId int64, uid int64) (res sql.Result, err error) {
	id, err := mysql.GetFollowUserById(followId, uid)
	if err != nil {
		return nil, err
	}
	if len(id) == 1 {
		res, err = mysql.DeleteFollowUserById(followId, uid)
	} else {
		err = errors.New("已经关注此用户！")
	}
	return
}
