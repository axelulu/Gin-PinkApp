package logic

import (
	"database/sql"
	"pinkacg/dao/mysql"
	"pinkacg/models"
)

// FollowStatus 关注状态
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

// FollowUser 关注用户
func FollowUser(followId int64, uid int64) (res sql.Result, err error) {
	id, err := mysql.GetFollowUserById(followId, uid)
	if err != nil {
		return nil, err
	}
	if len(id) == 0 {
		res, err = mysql.InsertFollowUserById(followId, uid)
	} else {
		err = mysql.ErrorUserFollowed
	}
	return
}

// UnFollowUser 取消关注
func UnFollowUser(followId int64, uid int64) (res sql.Result, err error) {
	id, err := mysql.GetFollowUserById(followId, uid)
	if err != nil {
		return nil, err
	}
	if len(id) == 1 {
		res, err = mysql.DeleteFollowUserById(followId, uid)
	} else {
		err = mysql.ErrorUserUnFollowed
	}
	return
}

// GetFansList 获取粉丝列表
func GetFansList(uid int64, isFans bool) (userMetas map[string]interface{}, err error) {
	userMetas = make(map[string]interface{}, 2)
	var userMeta []*models.FansMeta
	var users []*models.Follow
	if isFans {
		// 粉丝
		users, err = mysql.GetFollowsUserByFollowId(uid)
		if err != nil {
			return nil, err
		}
	} else {
		// 关注
		users, err = mysql.GetFollowsUserById(uid)
		if err != nil {
			return nil, err
		}
	}
	for _, userId := range users {
		// 根据作者id查询作者信息
		user := new(models.UserMeta)
		var follows []*models.Follow
		if isFans {
			// 粉丝
			user, err = mysql.GetUserById(userId.UserId)
			if err != nil {
				return nil, err
			}
			follows, err = mysql.GetFollowUserById(uid, userId.UserId)
			if err != nil {
				return nil, err
			}
		} else {
			// 关注
			user, err = mysql.GetUserById(userId.FollowId)
			if err != nil {
				return nil, err
			}
			follows, err = mysql.GetFollowUserById(userId.FollowId, uid)
			if err != nil {
				return nil, err
			}
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
