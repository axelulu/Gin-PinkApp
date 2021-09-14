package mysql

import "web_app/models"

func GetCategoryList(count int64) (category []*models.Category, err error) {
	sqlStr := `select category_slug, category_name from category limit ?`
	err = db.Select(&category, sqlStr, count)
	return
}
