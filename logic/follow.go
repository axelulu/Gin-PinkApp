package logic

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
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

func FollowStatus(followId int64, uid int64) (followStatus int, err error) {
	id, err := mysql.GetFollowUserById(followId, uid)
	if len(id) > 0 {
		// 已关注
		followStatus = 1
	} else if followId == uid {
		// 是本人
		followStatus = 2
	} else {
		// 没关注
		followStatus = 3
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

func GetFansList(uid int64, isFans bool) (userMetas map[string]interface{}, err error) {
	userMetas = make(map[string]interface{}, 2)
	var userMeta []*models.FansMeta
	var users []*models.Follow
	if isFans {
		// 粉丝
		users, err = mysql.GetFollowsUserByFollowId(uid)
	} else {
		// 关注
		users, err = mysql.GetFollowsUserById(uid)
	}
	for _, userId := range users {
		// 根据作者id查询作者信息
		user := new(models.UserMeta)
		var follows []*models.Follow
		if isFans {
			// 粉丝
			user, err = mysql.GetUserById(userId.UserId)
			follows, err = mysql.GetFollowUserById(userId.UserId, uid)
		} else {
			// 关注
			user, err = mysql.GetUserById(userId.FollowId)
			follows, err = mysql.GetFollowUserById(uid, userId.FollowId)
		}
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("post.AuthorID", userId.FollowId), zap.Error(err))
			continue
		}
		var meta *models.FansMeta
		if len(follows) > 0 {
			meta = &models.FansMeta{
				IsFollow: true,
				UserMeta: user,
			}
		} else {
			meta = &models.FansMeta{
				IsFollow: false,
				UserMeta: user,
			}
		}
		userMeta = append(userMeta, meta)
	}
	userMetas["list"] = userMeta
	userMetas["total"] = len(userMeta)
	return
}
