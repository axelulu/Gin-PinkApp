package oss

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"time"
	"web_app/controller"
)

func OSSUploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	fileExt := filepath.Ext(fileHeader.Filename)
	allowExts := []string{".jpg", ".png", ".gif", ".jpeg", ".doc", ".docx", ".ppt", ".pptx", ".xls", ".xlsx", ".pdf"}
	allowFlag := false
	for _, ext := range allowExts {
		if ext == fileExt {
			allowFlag = true
			break
		}
	}
	if !allowFlag {
		controller.ResponseError(c, controller.CodeServerBusy)
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
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	defer src.Close()

	res, err := OssUpload(fileKey, src)
	if err != nil {
		controller.ResponseError(c, controller.CodeServerBusy)
		return
	}
	controller.ResponseSuccess(c, res)
}
