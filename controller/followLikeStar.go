package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func FollowHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.FollowId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("PostCategoryListHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if fid == uid {
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.FollowUser(fid, uid)
	if err != nil {
		zap.L().Error("logic.FollowUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func UnFollowHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.FollowId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("PostCategoryListHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if fid == uid {
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.UnFollowUser(fid, uid)
	if err != nil {
		zap.L().Error("logic.UnFollowUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func LikeHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.LikeId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("LikeHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get strconv.ParseInt with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.LikePost(pid, uid)
	if err != nil {
		zap.L().Error("logic.LikePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func UnLikeHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.LikeId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("LikeHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get strconv.ParseInt with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.UnLikePost(pid, uid)
	if err != nil {
		zap.L().Error("logic.UnLikePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func StarHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.StarId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("StarHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get strconv.ParseInt with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.StarPost(pid, uid)
	if err != nil {
		zap.L().Error("logic.StarPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func UnStarHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.StarId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("UnStarHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get strconv.ParseInt with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.UnStarPost(pid, uid)
	if err != nil {
		zap.L().Error("logic.UnStarPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func CoinHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.CoinId)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("CoinHandle with invalid param", zap.Error(err))
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
		zap.L().Error("get strconv.ParseInt with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	coin, err := strconv.ParseInt(p.Coin, 10, 64)
	if err != nil {
		zap.L().Error("get strconv.ParseInt with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	msg, err := logic.CoinPost(pid, uid, coin)
	if err != nil {
		zap.L().Error("logic.CoinPost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}
