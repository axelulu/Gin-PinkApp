package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"pinkacg/dao/mysql"
	"pinkacg/logic"
	"pinkacg/models"
)

// CommentCreateHandle 增加评论
func CommentCreateHandle(c *gin.Context) {
	// 获取参数
	p := new(models.CommentCreate)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.CommentCreate with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID failed", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	user, err := logic.CommentCreate(uid, p)
	if err != nil {
		zap.L().Error("logic.CommentCreate failed", zap.Error(err))
		ResponseError(c, CodeCommentCreateFail)
		return
	}
	ResponseSuccess(c, user)
}

// CommentListHandle 获取评论列表
func CommentListHandle(c *gin.Context) {
	// 获取参数
	p := new(models.CommentList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.CommentList with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	user, err := logic.CommentList(p)
	if err != nil {
		zap.L().Error("logic.CommentList failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
