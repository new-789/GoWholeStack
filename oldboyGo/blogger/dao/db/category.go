package db

import (
	"github.com/GoWholeStack/oldboyGo/blogger/model"
	"github.com/jmoiron/sqlx"
)

// 分类数据操作 Demo

// InsertCategory 添加分类
func InsertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category(category_name, category_no) values(?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

// GetCategoryById 获取单个文章分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlStr := "select id, category_name, category_no from category where id=?"
	err = DB.Get(category, sqlStr, id)
	return
}

// GetCategoryList 获取多个文章分类
func GetCategoryList(categoryIds []int64) (categorylist []*model.Category, err error) {
	// 构建查询多个数据的 sql ,第一个返回值为 sql 语句，第二个返回值为传递给sql语句的参数切片 interface{} 类型
	sqlStr, args, err := sqlx.In("select id, category_name, category_no from category where id in(?)", categoryIds)
	if err != nil {
		return
	}
	// 查询
	err = DB.Select(&categorylist, sqlStr, args)
	return
}

// GetCategoryAll 获取所有文章分类
func GetCategoryAll() (categorylist []*model.Category, err error) {
	sqlStr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categorylist, sqlStr)
	return
}
