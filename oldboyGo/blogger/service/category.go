package service

import (
	"github.com/GoWholeStack/oldboyGo/blogger/dao/db"
	"github.com/GoWholeStack/oldboyGo/blogger/model"
)

// 主页分类功能实现

// GetAllCategoryList 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetCategoryAll()
	if err != nil {
		return
	}
	return
}