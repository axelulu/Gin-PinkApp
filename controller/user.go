package controller

import (
	"errors"
	"pinkacg/dao/mysql"
	"pinkacg/logic"
	"pinkacg/models"
	"strconv"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	// 获取参数
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("models.ParamSignUp with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 业务处理
	err := logic.SignUp(p)
	if err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		if errors.Is(err, mysql.ErrorValidateCode) {
			ResponseError(c, CodeValidateCode)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {
	// 获取请求参数及校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.ParamLogin with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 业务逻辑
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login.Login failed", zap.String("username", p.Email), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	ResponseSuccess(c, token)
}

func ForgetPwdHandler(c *gin.Context) {
	// 获取请求参数及校验
	p := new(models.UserForgetPwd)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("models.UserForgetPwd with invalid param", zap.Error(err))
		// 返回错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 业务逻辑
	res, err := logic.ForgetPwd(p)
	if err != nil {
		zap.L().Error("logic.ForgetPwd failed", zap.String("username", p.Email), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorValidateCode) {
			ResponseError(c, CodeValidateCode)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, res)
}

// UserHandle 获取用户信息
func UserHandle(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("id with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 逻辑处理
	user, err := logic.UserById(pid)
	if err != nil {
		zap.L().Error("logic.UserById failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// ProfileHandle 获取主页信息
func ProfileHandle(c *gin.Context) {
	// 获取用户ID
	uid, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID with invalid param", zap.Error(err))
		ResponseError(c, CodeCurrentUser)
		return
	}

	// 逻辑处理
	user, err := logic.UserById(uid)
	if err != nil {
		zap.L().Error("logic.UserById failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, user)
}

// UserCenterHandle 获取用户中心信息
func UserCenterHandle(c *gin.Context) {
	// 获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("id with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 逻辑处理
	user, err := logic.UserCenterById(pid)
	if err != nil {
		zap.L().Error("logic.UserCenterById failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// UserInfoUpdateHandle 更新用户信息
func UserInfoUpdateHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.UserUpdate)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.UserUpdate with invalid param", zap.Error(err))
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
	user, err := logic.UserInfoUpdate(uid, p)
	if err != nil {
		zap.L().Error("logic.UserInfoUpdate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func UserPasswordUpdateHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.UserPasswordUpdate)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.UserPasswordUpdate with invalid param", zap.Error(err))
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
	user, err := logic.UserPasswordUpdate(uid, p)
	if err != nil {
		zap.L().Error("logic.UserPasswordUpdate failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorValidateCode) {
			ResponseError(c, CodeValidateCode)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

func UserEmailUpdateHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.UserEmailUpdate)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("models.UserEmailUpdate with invalid param", zap.Error(err))
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
	user, err := logic.UserEmailUpdate(uid, p)
	if err != nil {
		zap.L().Error("logic.UserEmailUpdate failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorValidateCode) {
			ResponseError(c, CodeValidateCode)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}
