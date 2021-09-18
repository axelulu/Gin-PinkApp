package mysql

import "web_app/models"

func GetNewVersion() (version *models.Update, err error) {
	version = new(models.Update)
	sqlStr := `select version,url,meta from version where is_new=1`
	err = db.Get(version, sqlStr)
	return
}
