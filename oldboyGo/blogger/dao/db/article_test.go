package db

import (
	"github.com/GoWholeStack/oldboyGo/blogger/model"
	"testing"
	"time"
)

// 测试数据库操作Demo

func init() {
	// parseTime=true 将 mysql 中的时间类型字段解析为 go 结构体中的时间类型
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

// InsertArticle 插入文件操作函数测试
func TestInsertArticle(t *testing.T) {
	// 构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = `测试内容，测试内容，测试内容，测试内容,测试内容，测试内容，测试内容，
		测试内容,测试内容，测试内容，测试内容，测试内容,
		测试内容，测试内容，测试内容，测试内容,
		测试内容，测试内容，测试内容，测试内容
		`
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Username="三藏"
	article.ArticleInfo.Summary="测试内容..........."
	article.ArticleInfo.Title="测试"
	article.ArticleInfo.ViewCount=1
	articleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("articleId:%d\n", articleId)
}


// TestQueryArticleList 查询所有文件列表操作函数测试
func TestQueryArticleList(t *testing.T) {
	articleList, err := QueryArticleList(0, 15)
	if err != nil {
		panic(err)
	}
	t.Logf("%#v\n", len(articleList))
}

// TestQueryArticleDetail 测试获取单个文章操作函数
func TestQueryArticleDetail(t *testing.T) {
	articleDetail, err := QueryArticleDetail(1)
	if err != nil {
		panic(err)
	}
	t.Logf("reault:%#v\n", articleDetail)
}

// TestQueryArticleListByCategoryId 测试根据分类ID获取文章操作函数
func TestQueryArticleListByCategoryId(t *testing.T) {
	articleList, err := QueryArticleListByCategoryId(1, 0, 5)
	if err != nil {
		panic(err)
	}
	t.Logf("%#v\n", articleList)
}