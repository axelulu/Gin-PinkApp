package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func CommentCreateHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.CommentCreate)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("CommentPublishHandle with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.CommentCreate(uid, p)
	if err != nil {
		zap.L().Error("logic.UserMetaById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func CommentListHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.CommentList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("CommentPublishHandle with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	user, err := logic.CommentList(p)
	if err != nil {
		zap.L().Error("logic.UserMetaById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
