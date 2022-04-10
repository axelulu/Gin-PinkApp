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

// CategoryListHandle 获取所有分类
func CategoryListHandle(c *gin.Context) {
	// 获取参数
	p := new(models.CategoryList)
	if err := c.ShouldBindQuery(&p); err != nil {
		// 记录日志
		zap.L().Error("models.CategoryList with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//逻辑处理
	category, err := logic.CategoryList(p)
	if err != nil {
		zap.L().Error("logic.CategoryList failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorCatEmpty) {
			ResponseError(c, CodeCatEmpty)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, category)
}
