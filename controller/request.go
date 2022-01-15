package controller

import (
	"github.com/gin-gonic/gin"
	"pinkacg/dao/mysql"
)

const CtxUserIDKey = "userID"

// 获取当前用户
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = mysql.ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = mysql.ErrorUserNotLogin
		return
	}
	return
}
