package service

import (
	"github.com/GoWholeStack/oldboyGo/blogger/dao/db"
	"github.com/GoWholeStack/oldboyGo/blogger/model"
)

// 主页文章业务逻辑层实现

// GetArticleRecordList 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.QueryArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	// 校验获取到的数据
	if len(articleInfoList) <= 0 {
		return
	}
	// 2. 获取文章对应的分类
	categoryIds := getCategoryIds(articleInfoList)
	// 2.1：获取所有分类信息
	categorylist, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 3. 返回页面做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		// 根据当前文章，生成 ArticleRecord 数据库对应的结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 从文章中取出分类 id
		categoryId := article.CategoryId
		// 遍历所有分类信息列表
		for _, category := range categorylist {
			/* 判断每篇文章的分类id，如果等于分类列表中的 id，则说明获取成功，
			   给结构体中的 Category 进行赋值操作，这样每篇文章和对应的分类信息
			   的对应情况就不会出错了
			*/
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

// GetCategoryIds 根据多个文章的id，获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LAB:
	// 遍历文章，得到每个文章，在根据每个文章的id 获取它的分类 id
	for _, article := range articleInfoList {
		// 从当前文章id，取出分类 id
		categoryId := article.CategoryId
		// 去重，防止重复
		for _, id := range ids {
			// 看当前的id是否存储在
			if id == categoryId {
				continue LAB
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// GetArticleRecordListById 根据分类 ID 获取该类文章和它们对应的分类信息
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 1.获取文章列表
	articleInfoList, err := db.QueryArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	// 校验获取到的数据
	if len(articleInfoList) <= 0 {
		return
	}
	// 2. 获取文章对应的分类
	categoryIds := getCategoryIds(articleInfoList)
	// 2.1：获取所有分类信息
	categorylist, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	// 3. 返回页面做聚合
	// 遍历所有文章
	for _, article := range articleInfoList {
		// 根据当前文章，生成 ArticleRecord 数据库对应的结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		// 从文章中取出分类 id
		categoryId := article.CategoryId
		// 遍历所有分类信息列表
		for _, category := range categorylist {
			/* 判断每篇文章的分类id，如果等于分类列表中的 id，则说明获取成功，
			   给结构体中的 Category 进行赋值操作，这样每篇文章和对应的分类信息
			   的对应情况就不会出错了
			*/
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}
