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

// ContactListHandle 获取联系人列表
func ContactListHandle(c *gin.Context) {
	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	user, err := logic.GetContactList(uid)
	if err != nil {
		zap.L().Error("logic.GetContactList failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		if errors.Is(err, mysql.ErrorUserChat) {
			ResponseError(c, CodeUserChat)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// ContactItemHandle 获取单个联系人信息
func ContactItemHandle(c *gin.Context) {
	// 1. 获取参数
	sidStr := c.Param("id")
	sid, err := strconv.ParseInt(sidStr, 10, 64)
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

	user, err := logic.GetContactItem(uid, sid)
	if err != nil {
		zap.L().Error("logic.GetContactListByUserId failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		if errors.Is(err, mysql.ErrorUserChat) {
			ResponseError(c, CodeUserChat)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// ChatListHandle 聊天信息列表
func ChatListHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.ChatList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.ChatList with invalid param", zap.Error(err))
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
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	user, err := logic.GetChatList(uid, p)
	if err != nil {
		zap.L().Error("logic.GetChatList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// ContactCreateHandle 创建对话
func ContactCreateHandle(c *gin.Context) {
	// 1. 获取请求参数及校验
	contactAdd := new(models.ContactAdd)
	if err := c.ShouldBindJSON(&contactAdd); err != nil {
		// 记录日志
		zap.L().Error("models.ContactAdd with invalid param", zap.Error(err))
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

	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	user, err := logic.AddContactList(uid, sid)
	if err != nil {
		zap.L().Error("logic.AddContactList failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		if errors.Is(err, mysql.ErrorContactExist) {
			ResponseError(c, CodeContactExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
