package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/utils"
)

// UploadImage 上传图片到OSS
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

	// 文件大小限制已移除

	// 打开文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "打开文件失败"})
		return
	}
	defer src.Close()

	// 上传到OSS
	fileURL, err := utils.UploadToOSS(src, file.Filename, "image")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传图片失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "图片上传成功",
		"url":      fileURL,
		"filename": file.Filename,
	})
}

// UploadVideo 上传视频到OSS
func UploadVideo(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请选择要上传的视频"})
		return
	}

	// 验证文件类型
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".mp4":  true,
		".avi":  true,
		".mov":  true,
		".wmv":  true,
		".flv":  true,
		".mkv":  true,
		".webm": true,
	}

	if !allowedExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的视频格式，请上传 mp4、avi、mov、wmv、flv、mkv 或 webm 格式"})
		return
	}

	// 文件大小限制已移除

	// 打开文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "打开文件失败"})
		return
	}
	defer src.Close()

	// 上传到OSS
	fileURL, err := utils.UploadToOSS(src, file.Filename, "video")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传视频失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "视频上传成功",
		"url":      fileURL,
		"filename": file.Filename,
	})
}

// DeleteImage 删除OSS中的图片
func DeleteImage(c *gin.Context) {
	type DeleteInput struct {
		URL string `json:"url" binding:"required"`
	}

	var input DeleteInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 从OSS删除文件
	if err := utils.DeleteFromOSS(input.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文件失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
