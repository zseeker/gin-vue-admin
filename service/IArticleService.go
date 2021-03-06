package service

import (
	"github.com/bingjian-zhu/gin-vue-admin/models"
)

//IArticleService Article接口定义
type IArticleService interface {
	// //GetArticle 根据id获取Article
	// GetArticle(id int) models.Article
	//GetArticles 分页返回文章
	GetArticles(page, pagesize int, maps map[string]interface{}) *[]models.Article
	// //AddArticle 新增Article
	// AddArticle(article models.Article) bool
	// //ExistArticleByID 根据ID判断Article是否存在
	// ExistArticleByID(id int) bool
	// //EditArticle 编辑Article
	// EditArticle(article models.Article) bool
	// //DeleteArticle 删除Article
	// DeleteArticle(id int) bool
	// //GetArticles 分页返回Articles(v2)
	// GetArticlesv2(PageNum int, PageSize int, maps interface{}) (article []models.Article)
}
