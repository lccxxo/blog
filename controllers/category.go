package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/models"
)

type CategoryInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := database.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "分类名称已存在"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "分类创建成功",
		"category": category,
	})
}

// GetCategories 获取所有分类
func GetCategories(c *gin.Context) {
	var categories []models.Category
	database.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GetCategory 获取单个分类
func GetCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	var input CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = input.Name
	category.Description = input.Description

	database.DB.Save(&category)

	c.JSON(http.StatusOK, gin.H{
		"message":  "分类更新成功",
		"category": category,
	})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "分类不存在"})
		return
	}

	database.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"message": "分类删除成功"})
}




