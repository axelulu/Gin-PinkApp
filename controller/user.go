package controller

import (
	"errors"
	"strconv"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数及校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 2. 业务逻辑
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, token)
}

func UserHandle(c *gin.Context) {
	// 1. 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.UserById(pid)
	if err != nil {
		zap.L().Error("logic.UserById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func ProfileHandle(c *gin.Context) {
	// 1. 获取参数
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.UserById(uid)
	if err != nil {
		zap.L().Error("logic.UserById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func UserCenterHandle(c *gin.Context) {
	// 1. 获取参数
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("get getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	user, err := logic.UserCenterById(uid)
	if err != nil {
		zap.L().Error("logic.UserMetaById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
