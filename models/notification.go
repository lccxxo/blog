package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 通知接收者
	UserID uint `gorm:"not null;index" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// 通知类型：comment_reply（评论回复）、mention（@提及）
	Type string `gorm:"type:varchar(50);not null" json:"type"`

	// 通知内容
	Content string `gorm:"type:varchar(500)" json:"content"`

	// 相关的评论ID
	CommentID *uint    `gorm:"index" json:"comment_id,omitempty"`
	Comment   *Comment `gorm:"foreignKey:CommentID" json:"comment,omitempty"`

	// 相关的文章ID
	ArticleID *uint    `gorm:"index" json:"article_id,omitempty"`
	Article   *Article `gorm:"foreignKey:ArticleID" json:"article,omitempty"`

	// 触发通知的用户（谁回复了你）
	FromUserID uint `gorm:"not null" json:"from_user_id"`
	FromUser   User `gorm:"foreignKey:FromUserID" json:"from_user"`

	// 是否已读
	IsRead bool `gorm:"default:false;index" json:"is_read"`
}
