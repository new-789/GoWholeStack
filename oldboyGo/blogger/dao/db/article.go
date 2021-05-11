package db

// 文章数据库操作 Demo

import (
	"github.com/GoWholeStack/oldboyGo/blogger/model"
	_ "github.com/go-sql-driver/mysql"
)

// InsertArticle 插入文章内容到数据库
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	// 加验证，如果是结构体中的数据为空则直接返回
	if article == nil {
		return
	}
	sqlStr := `insert into 
    				article(category_id,content,title,username,summary,view_count,comment_count)
			   values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, article.ArticleInfo.CategoryId, article.Content, article.Title,
		article.Username, article.Summary, article.ViewCount, article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

// QueryArticleList 查询所有文章列表操作内容，做个分页
func QueryArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	// 验证分页参数
	if pageNum < 0 || pageSize < 0 {
		return
	}
	// 做数据查询，并按时间按降序排序
	sqlStr := `select 
				  id,category_id,title,username,summary,view_count,comment_count,create_time
				from 
				  article
				where 
				  status=1
				order by 
				  create_time desc limit ?,?` // limit 用于分页
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)

	if err != nil {
		return
	}
	return
}

// QueryArticleDetail 根据文章ID查询单个文章
func QueryArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	// 验证传入的文章Id
	if articleId < 0 {
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlStr := `select
					id,category_id,title,username,content,summary,view_count,comment_count,create_time
			   from 
					article
			   where 
			        id=? and status=1
   			   `
	err = DB.Get(&articleDetail, sqlStr, articleId)
	if err != nil {
		return
	}
	return
}

// QueryArticleListByCategoryId 根据分类 ID 查询这一类的文章
func QueryArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	// 参数验证
	if categoryId < 0 || pageNum < 0 || pageSize < 0 {
		return
	}
	sqlStr := `select
			  		id,category_id,title,username,summary,view_count,comment_count,create_time
			   from
					article
			   where 
					category_id = ? and status=1
			   order by
					create_time desc limit ?,?
			   `
	err = DB.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	return
}
