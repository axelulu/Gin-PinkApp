package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
)

func CategoryListHandle(c *gin.Context) {
	//参数校验
	p := new(models.CategoryList)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 记录日志
		zap.L().Error("SignUp with invalid param", zap.Error(err))
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
		zap.L().Error("logic.CategoryLis failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, category)
}

func AddCategoryHandle(c *gin.Context) {

}
