package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func HomeHandle(c *gin.Context) {
	p := new(models.Home)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("HomeHandle with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	home, err := logic.HomeList(p)
	if err != nil {
		zap.L().Error("logic.HomeList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, home)
}

func PostCategoryListHandle(c *gin.Context) {
	p := new(models.PostCategoryList)
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

	post, err := logic.PostCategoryList(p)
	if err != nil {
		zap.L().Error("logic.PostCategoryList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

func PostByIdHandle(c *gin.Context) {
	// 1. 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	post, err := logic.PostById(pid)
	if err != nil {
		zap.L().Error("logic.PostById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

func RankingHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.PostRankingList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("RankingHandle with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	post, err := logic.PostRanking(p)
	if err != nil {
		zap.L().Error("logic.PostRanking failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}
