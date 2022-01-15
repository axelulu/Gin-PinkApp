package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"pinkacg/dao/mysql"
	"pinkacg/logic"
	"pinkacg/models"
	"strconv"
)

// FollowStatusHandle 获取关注状态
func FollowStatusHandle(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	fid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("id with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.FollowStatus(fid, uid)
	if err != nil {
		zap.L().Error("logic.FollowStatus failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// FollowHandle 关注用户
func FollowHandle(c *gin.Context) {
	// 获取参数
	p := new(models.FollowId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.FollowId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fid, err := strconv.ParseInt(p.FollowId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	if fid == uid {
		ResponseError(c, CodeIsCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.FollowUser(fid, uid)
	if err != nil {
		zap.L().Error("logic.FollowUser failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserFollowed) {
			ResponseError(c, CodeUserFollowed)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// UnFollowHandle 取消关注
func UnFollowHandle(c *gin.Context) {
	// 获取参数
	p := new(models.FollowId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.FollowId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fid, err := strconv.ParseInt(p.FollowId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	if fid == uid {
		ResponseError(c, CodeIsCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.UnFollowUser(fid, uid)
	if err != nil {
		zap.L().Error("logic.UnFollowUser failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserUnFollowed) {
			ResponseError(c, CodeUserUnFollowed)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// LikeHandle 喜欢
func LikeHandle(c *gin.Context) {
	// 获取参数
	p := new(models.LikeId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.LikeId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	pid, err := strconv.ParseInt(p.PostId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.LikePost(pid, uid)
	if err != nil {
		zap.L().Error("logic.LikePost failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostLiked) {
			ResponseError(c, CodePostLiked)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// UnLikeHandle 不喜欢
func UnLikeHandle(c *gin.Context) {
	// 获取参数
	p := new(models.LikeId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.LikeId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	pid, err := strconv.ParseInt(p.PostId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.UnLikePost(pid, uid)
	if err != nil {
		zap.L().Error("logic.UnLikePost failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostUnLiked) {
			ResponseError(c, CodePostUnLiked)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// StarHandle 收藏
func StarHandle(c *gin.Context) {
	// 获取参数
	p := new(models.StarId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.StarId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	pid, err := strconv.ParseInt(p.PostId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.StarPost(pid, uid)
	if err != nil {
		zap.L().Error("logic.StarPost failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostStared) {
			ResponseError(c, CodePostStared)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// UnStarHandle 取消收藏
func UnStarHandle(c *gin.Context) {
	// 获取参数
	p := new(models.StarId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.StarId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	pid, err := strconv.ParseInt(p.PostId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.UnStarPost(pid, uid)
	if err != nil {
		zap.L().Error("logic.UnStarPost failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostUnStared) {
			ResponseError(c, CodePostUnStared)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// CoinHandle 投币
func CoinHandle(c *gin.Context) {
	// 获取参数
	p := new(models.CoinId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.CoinId with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	pid, err := strconv.ParseInt(p.PostId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	coin, err := strconv.ParseInt(p.Coin, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.CoinPost(pid, uid, coin)
	if err != nil {
		zap.L().Error("logic.CoinPost failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostCoined) {
			ResponseError(c, CodePostCoined)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// FollowListHandle 获取关注列表
func FollowListHandle(c *gin.Context) {
	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.GetFansList(uid, false)
	if err != nil {
		zap.L().Error("logic.GetFollowList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

// FansListHandle 获取粉丝列表
func FansListHandle(c *gin.Context) {
	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	msg, err := logic.GetFansList(uid, true)
	if err != nil {
		zap.L().Error("logic.GetFansList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}
