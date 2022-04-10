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

// SearchHandle 获取搜索列表
func SearchHandle(c *gin.Context) {
	// 获取参数
	p := new(models.Search)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.Search with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 逻辑处理
	post, err := logic.Search(p)
	if err != nil {
		zap.L().Error("logic.Search failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorPostMeta) {
			ResponseError(c, CodePostMeta)
			return
		}
		if errors.Is(err, mysql.ErrorUserMeta) {
			ResponseError(c, CodeUserMeta)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}
