package mysql

import "pinkacg/models"

// GetCategoryList 从数据库中获取所有分类信息
func GetCategoryList(count int64) (category []*models.Category, err error) {
	sqlStr := `select category_id, category_name from categories limit ?`
	err = db.Select(&category, sqlStr, count)
	return
}
