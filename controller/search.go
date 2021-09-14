package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func SearchHandle(c *gin.Context) {
	// 1. 获取参数
	p := new(models.Search)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("SearchHandle with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	post, err := logic.Search(p)
	print(post)
	if err != nil {
		zap.L().Error("logic.Search failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}
