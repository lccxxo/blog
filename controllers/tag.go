package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/models"
)

type TagInput struct {
	Name string `json:"name" binding:"required"`
}

// CreateTag 创建标签
func CreateTag(c *gin.Context) {
	var input TagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag := models.Tag{
		Name: input.Name,
	}

	if err := database.DB.Create(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标签名称已存在"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "标签创建成功",
		"tag":     tag,
	})
}

// GetTags 获取所有标签
func GetTags(c *gin.Context) {
	var tags []models.Tag
	database.DB.Find(&tags)

	c.JSON(http.StatusOK, gin.H{"tags": tags})
}

// GetTag 获取单个标签
func GetTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tag": tag})
}

// UpdateTag 更新标签
func UpdateTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	var input TagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.Name = input.Name

	database.DB.Save(&tag)

	c.JSON(http.StatusOK, gin.H{
		"message": "标签更新成功",
		"tag":     tag,
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	id := c.Param("id")

	var tag models.Tag
	if err := database.DB.First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "标签不存在"})
		return
	}

	database.DB.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"message": "标签删除成功"})
}




