package logic

import (
	"pinkacg/dao/mysql"
	"pinkacg/models"
)

// CategoryList 获取分类列表
func CategoryList(p *models.CategoryList) (category []*models.Category, err error) {
	category, err = mysql.GetCategoryList(p.Size)
	if len(category) == 0 {
		err = mysql.ErrorCatEmpty
	}
	return
}
