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

// SendRegEmailHandle 发送用户注册邮件
func SendRegEmailHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.Email)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Email with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	err := logic.UserRegSendEmail(p, Captcha(6))
	if err != nil {
		zap.L().Error("logic.UserRegSendEmail failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorEmailExist) {
			ResponseError(c, CodeEmailExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// SendForgetPwdEmailHandle 发送用户忘记密码邮件
func SendForgetPwdEmailHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.Email)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Email with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	err := logic.UserForgetPwdSendEmail(p, Captcha(6))
	if err != nil {
		zap.L().Error("logic.UserRegSendEmail failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// SendChangePwdEmailHandle 发送用户修改密码邮件
func SendChangePwdEmailHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.Email)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Email with invalid param", zap.Error(err))
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
	err = logic.UserChangePwdSendEmail(p, Captcha(6), uid)
	if err != nil {
		zap.L().Error("logic.UserChangePwdSendEmail failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// SendChangeEmailHandle 发送用户修改邮箱邮件
func SendChangeEmailHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.Email)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Email with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	err := logic.UserChangeEmailSendEmail(p, Captcha(6))
	if err != nil {
		zap.L().Error("logic.UserChangePwdSendEmail failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorEmailExist) {
			ResponseError(c, CodeEmailExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
