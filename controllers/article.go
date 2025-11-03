package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/models"
)

type ArticleInput struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Summary     string `json:"summary"`
	CoverImage  string `json:"cover_image"`
	CategoryID  uint   `json:"category_id"`
	TagIDs      []uint `json:"tag_ids"`
	IsPublished bool   `json:"is_published"`
}

// CreateArticle 创建文章
func CreateArticle(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
		Title:       input.Title,
		Content:     input.Content,
		Summary:     input.Summary,
		CoverImage:  input.CoverImage,
		AuthorID:    userID.(uint),
		CategoryID:  input.CategoryID,
		IsPublished: input.IsPublished,
	}

	// 处理标签
	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		database.DB.Where("id IN ?", input.TagIDs).Find(&tags)
		article.Tags = tags
	}

	if err := database.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	// 加载关联数据
	database.DB.Preload("Author").Preload("Category").Preload("Tags").First(&article, article.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "文章创建成功",
		"article": article,
	})
}

// GetArticles 获取文章列表
func GetArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	categoryID := c.Query("category_id")
	tagID := c.Query("tag_id")

	var articles []models.Article
	query := database.DB.Preload("Author").Preload("Category").Preload("Tags")

	// 筛选条件
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if tagID != "" {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Where("article_tags.tag_id = ?", tagID)
	}

	// 只显示已发布的文章（公开接口）
	query = query.Where("is_published = ?", true)

	var total int64
	query.Model(&models.Article{}).Count(&total)

	offset := (page - 1) * pageSize
	query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles)

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

// GetArticle 获取单篇文章
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	if err := database.DB.Preload("Author").Preload("Category").Preload("Tags").First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 增加浏览量
	database.DB.Model(&article).Update("view_count", article.ViewCount+1)

	c.JSON(http.StatusOK, gin.H{"article": article})
}

// UpdateArticle 更新文章
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 验证作者权限
	if article.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改此文章"})
		return
	}

	var input ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新文章字段
	article.Title = input.Title
	article.Content = input.Content
	article.Summary = input.Summary
	article.CoverImage = input.CoverImage
	article.CategoryID = input.CategoryID
	article.IsPublished = input.IsPublished

	// 更新标签
	if len(input.TagIDs) > 0 {
		var tags []models.Tag
		database.DB.Where("id IN ?", input.TagIDs).Find(&tags)
		database.DB.Model(&article).Association("Tags").Replace(tags)
	}

	database.DB.Save(&article)

	// 重新加载关联数据
	database.DB.Preload("Author").Preload("Category").Preload("Tags").First(&article, article.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "文章更新成功",
		"article": article,
	})
}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var article models.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 验证作者权限
	if article.AuthorID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除此文章"})
		return
	}

	database.DB.Delete(&article)

	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
}

// GetMyArticles 获取我的文章列表
func GetMyArticles(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var articles []models.Article
	query := database.DB.Where("author_id = ?", userID).
		Preload("Category").Preload("Tags")

	var total int64
	query.Model(&models.Article{}).Count(&total)

	offset := (page - 1) * pageSize
	query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles)

	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}




