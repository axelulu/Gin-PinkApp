package mysql

import "pinkacg/models"

func GetNewVersion() (version *models.Update, err error) {
	version = new(models.Update)
	sqlStr := `select version,url,meta from versions where is_new=1`
	err = db.Get(version, sqlStr)
	return
}
