package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
)

func UpdateHandle(c *gin.Context) {
	version, err := logic.GetUpdate()
	if err != nil {
		zap.L().Error("logic.GetUpdate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, version)
}
