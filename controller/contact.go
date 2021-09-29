package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func ContactListHandle(c *gin.Context) {
	// 1. 获取参数
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.GetContactList(uid)
	if err != nil {
		zap.L().Error("logic.GetContactListByUserId failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func ChatListHandle(c *gin.Context) {
	// 1. 获取参数
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.GetChatList(uid)
	if err != nil {
		zap.L().Error("logic.GetChatList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func ContactPublishHandle(c *gin.Context) {
	// 1. 获取请求参数及校验
	contactAdd := new(models.ContactAdd)
	if err := c.ShouldBindJSON(&contactAdd); err != nil {
		// 记录日志
		zap.L().Error("ContactPublishHandle with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	sid, err := strconv.ParseInt(contactAdd.SendId, 10, 64)

	// 1. 获取参数
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.AddContactList(uid, sid)
	if err != nil {
		zap.L().Error("logic.AddContactList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
