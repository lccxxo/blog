package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/models"
)

type CommentInput struct {
	Content   string `json:"content" binding:"required"`
	ArticleID uint   `json:"article_id" binding:"required"`
	ParentID  *uint  `json:"parent_id"` // 可选，用于回复评论
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证文章是否存在
	var article models.Article
	if err := database.DB.First(&article, input.ArticleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	// 如果是回复评论，验证父评论是否存在
	if input.ParentID != nil {
		var parentComment models.Comment
		if err := database.DB.First(&parentComment, *input.ParentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "父评论不存在"})
			return
		}
		// 确保父评论属于同一篇文章
		if parentComment.ArticleID != input.ArticleID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "父评论不属于该文章"})
			return
		}
	}

	comment := models.Comment{
		Content:   input.Content,
		ArticleID: input.ArticleID,
		UserID:    userID.(uint),
		ParentID:  input.ParentID,
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}

	// 加载用户信息
	database.DB.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "评论发表成功",
		"comment": comment,
	})
}

// GetArticleComments 获取文章的评论列表
func GetArticleComments(c *gin.Context) {
	articleID := c.Param("id")

	// 验证文章是否存在
	var article models.Article
	if err := database.DB.First(&article, articleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	var comments []models.Comment
	// 只获取顶层评论（没有父评论的）
	database.DB.Where("article_id = ? AND parent_id IS NULL", articleID).
		Preload("User").
		Preload("Replies.User").
		Order("created_at DESC").
		Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"total":    len(comments),
	})
}

// GetComment 获取单个评论
func GetComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := database.DB.Preload("User").Preload("Replies.User").First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

// UpdateComment 更新评论
func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}

	// 验证是否是评论作者
	if comment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限修改此评论"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = input.Content
	database.DB.Save(&comment)

	// 重新加载用户信息
	database.DB.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "评论更新成功",
		"comment": comment,
	})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "评论不存在"})
		return
	}

	// 验证是否是评论作者
	if comment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "无权限删除此评论"})
		return
	}

	// 删除评论及其所有回复
	database.DB.Where("parent_id = ?", comment.ID).Delete(&models.Comment{})
	database.DB.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}

// GetMyComments 获取我的评论列表
func GetMyComments(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var comments []models.Comment
	query := database.DB.Where("user_id = ?", userID).
		Preload("Article").
		Preload("User")

	var total int64
	query.Model(&models.Comment{}).Count(&total)

	offset := (page - 1) * pageSize
	query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
