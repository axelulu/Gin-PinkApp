package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"pinkacg/pkg/oss"
	"pinkacg/settings"
	"time"
)

// UploadHandle 上传文件
func UploadHandle(c *gin.Context) {
	// 获取参数
	fileHeader, err := c.FormFile("file")
	if err != nil {
		zap.L().Error("c.FormFile failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 文件扩展名
	fileExt := filepath.Ext(fileHeader.Filename)
	// 允许的文件类型
	allowExts := []string{".mp4", ".jpg", ".png", ".gif", ".jpeg"}
	allowFlag := false
	for _, ext := range allowExts {
		if ext == fileExt {
			allowFlag = true
			break
		}
	}
	if !allowFlag {
		zap.L().Error("allowFlag failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//文件存放路径
	now := time.Now()
	var fileDir string
	if fileExt == ".mp4" {
		if fileHeader.Size > 1024*1024*settings.Conf.VideoSize {
			zap.L().Error("fileHeader.Size over", zap.Error(err))
			ResponseError(c, CodeUploadFileOver)
			return
		}
		fileDir = fmt.Sprintf("PinkAcg/video/%s", now.Format("2006-01-01"))
	} else {
		if fileHeader.Size > 1024*1024*settings.Conf.PicSize {
			zap.L().Error("fileHeader.Size over", zap.Error(err))
			ResponseError(c, CodeUploadFileOver)
			return
		}
		fileDir = fmt.Sprintf("PinkAcg/pic/%s", now.Format("2006-01-01"))
	}

	//文件名称
	timeStamp := now.Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	// 文件key
	fileKey := filepath.Join(fileDir, fileName)

	src, err := fileHeader.Open()
	if err != nil {
		zap.L().Error("fileHeader.Open failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	defer src.Close()

	// 开始上传
	res, err := oss.OssUpload(fileKey, src)
	if err != nil {
		zap.L().Error("oss.OssUpload failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, res)
}
