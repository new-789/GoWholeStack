package model

import "time"

// ArticleInfo 定义文章结构体
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
	Summary      string    `db:"summary"`
	CreateTime   time.Time `db:"create_time"`
}

// ArticleDetail 用于文章详情也实体
// 之所以单独将文章详情页单独定义一个结构体，是为了加载首页时不需要该字段，所以便不加载该字段从而提升效率
type ArticleDetail struct {
	ArticleInfo
	// 文章内容
	Content string `db:"content"`
	// 在文章详情也中用于关联文章分类
	Category
}

// ArticleRecord 用于在文章详情也上下关联
type ArticleRecord struct {
	ArticleInfo
	Category
}
