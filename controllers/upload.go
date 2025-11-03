package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadImage 上传图片
func UploadImage(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请选择要上传的图片"})
		return
	}

	// 验证文件类型
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
		".svg":  true,
	}

	if !allowedExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的图片格式，请上传 jpg、png、gif、webp 或 svg 格式"})
		return
	}

	// 验证文件大小（限制为5MB）
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "图片大小不能超过5MB"})
		return
	}

	// 创建上传目录
	uploadDir := "uploads/images"
	// 按日期分类存储
	dateDir := time.Now().Format("2006/01/02")
	fullDir := filepath.Join(uploadDir, dateDir)

	if err := os.MkdirAll(fullDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建上传目录失败"})
		return
	}

	// 生成唯一文件名
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	filePath := filepath.Join(fullDir, filename)

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存图片失败"})
		return
	}

	// 返回文件URL（使用正斜杠，适用于URL）
	fileURL := fmt.Sprintf("/uploads/images/%s/%s", strings.ReplaceAll(dateDir, "\\", "/"), filename)

	c.JSON(http.StatusOK, gin.H{
		"message":  "图片上传成功",
		"url":      fileURL,
		"filename": file.Filename,
	})
}

// DeleteImage 删除图片
func DeleteImage(c *gin.Context) {
	type DeleteInput struct {
		URL string `json:"url" binding:"required"`
	}

	var input DeleteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 验证URL格式
	if !strings.HasPrefix(input.URL, "/uploads/images/") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的图片URL"})
		return
	}

	// 删除文件
	filePath := "." + input.URL
	if err := os.Remove(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除图片失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "图片删除成功"})
}


