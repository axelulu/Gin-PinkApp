package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func CategoryList(p *models.CategoryList) (category []*models.Category, err error) {
	category, err = mysql.GetCategoryList(p.Size)
	return
}
