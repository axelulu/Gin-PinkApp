package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetUpdate() (version *models.Update, err error) {
	version, err = mysql.GetNewVersion()
	return
}
