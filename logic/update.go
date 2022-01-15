package logic

import (
	"pinkacg/dao/mysql"
	"pinkacg/models"
)

func GetUpdate() (version *models.Update, err error) {
	version, err = mysql.GetNewVersion()
	return
}
