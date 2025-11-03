package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Content   string         `gorm:"type:text;not null" json:"content" binding:"required"`
	ArticleID uint           `gorm:"not null;index" json:"article_id"`
	Article   *Article       `gorm:"foreignKey:ArticleID" json:"article,omitempty"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	ParentID  *uint          `gorm:"index" json:"parent_id"` // 父评论ID，用于回复
	Replies   []Comment      `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
}




