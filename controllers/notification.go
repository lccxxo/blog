package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/models"
)

// GetNotifications 获取当前用户的通知列表
func GetNotifications(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	unreadOnly := c.Query("unread_only") == "true"

	var notifications []models.Notification
	query := database.DB.Model(&models.Notification{}).Where("user_id = ?", userID).
		Preload("FromUser").
		Preload("Article").
		Preload("Comment.User")

	if unreadOnly {
		query = query.Where("is_read = ?", false)
	}

	var total int64
	query.Count(&total)

	offset := (page - 1) * pageSize
	query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&notifications)

	c.JSON(http.StatusOK, gin.H{
		"notifications": notifications,
		"total":         total,
		"page":          page,
		"pageSize":      pageSize,
	})
}

// GetUnreadCount 获取未读通知数量
func GetUnreadCount(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var count int64
	database.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Count(&count)

	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}

// MarkAsRead 标记通知为已读
func MarkAsRead(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var notification models.Notification
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "通知不存在"})
		return
	}

	notification.IsRead = true
	database.DB.Save(&notification)

	c.JSON(http.StatusOK, gin.H{
		"message": "已标记为已读",
	})
}

// MarkAllAsRead 标记所有通知为已读
func MarkAllAsRead(c *gin.Context) {
	userID, _ := c.Get("user_id")

	database.DB.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true)

	c.JSON(http.StatusOK, gin.H{
		"message": "所有通知已标记为已读",
	})
}

// DeleteNotification 删除通知
func DeleteNotification(c *gin.Context) {
	userID, _ := c.Get("user_id")
	id := c.Param("id")

	var notification models.Notification
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "通知不存在"})
		return
	}

	database.DB.Delete(&notification)

	c.JSON(http.StatusOK, gin.H{
		"message": "通知已删除",
	})
}
