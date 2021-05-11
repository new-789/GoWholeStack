package db

import "testing"

// 测试数据库操作Demo

func init() {
	// parseTime=true 将 mysql 中的时间类型字段解析为 go 结构体中的时间类型
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// TestGetCategoryById 测试从数据库获取单个分类
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(0)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

// TestGetCategoryList 测试从数据库获取多个分类
func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1,2,3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList{
		t.Logf("id:%d category:%v\n", v.CategoryId, v)
	}
}

// TestGetCategoryAll 测试从数据库获取所有分类
func TestGetCategoryAll(t *testing.T) {
	categorylist, err := GetCategoryAll()
	if err != nil {
		panic(err)
	}
	for _, v := range categorylist {
		t.Logf("id:%d caregoryName:%s categoryNo:%v\n", v.CategoryId, v.CategoryName, v.CategoryNo)
	}
}