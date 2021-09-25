package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"time"
	"web_app/pkg/oss"
)

func UploadHandle(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	allowExts := []string{".mp4", ".jpg", ".png", ".gif", ".jpeg", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx", ".pdf"}
	allowFlag := false
	for _, ext := range allowExts {
		if ext == fileExt {
			allowFlag = true
			break
		}
	}
	if !allowFlag {
		ResponseError(c, CodeServerBusy)
		return
	}

	now := time.Now()
	//文件存放路径
	fileDir := fmt.Sprintf("articles/%s", now.Format("200601"))

	//文件名称
	timeStamp := now.Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	// 文件key
	fileKey := filepath.Join(fileDir, fileName)

	src, err := fileHeader.Open()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	defer src.Close()

	res, err := oss.OssUpload(fileKey, src)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, res)
}
