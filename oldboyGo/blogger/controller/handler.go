package controller

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/blogger/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// IndexHandle 主页访问控制器
func IndexHandle(c *gin.Context) {
	// 从 service 模块获取数据

	// 1. 加载文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	fmt.Println("---------------------------------->", articleRecordList)
	if err != nil {
		fmt.Println("service getArticleRecoredList failed err", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 2. 加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}


	/* gin.H 本质上的一个 map，如下伪代码示例
		var data map[string]interface{} = make(map[string]interface{}, 16)
		data["article_list"] = articleList
		data["category_list"] = categoryList
		c.HTML(http.StatusOK, "views/index.html", data)
	*/
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list": articleRecordList,
		"category_list": categoryList,
	})
}

// CategoryList 点击分类云进行分类
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	// 将获取到的前段字符串 id 装为 int64 类型
	categoryId, err := strconv.ParseInt(categoryIdStr,10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 根据分类 id ，获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	// 再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCategoryList()
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list": articleRecordList,
		"category_list": categoryList,
	})
}