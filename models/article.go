package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"not null" json:"title" binding:"required"`
	Content     string         `gorm:"type:text;not null" json:"content" binding:"required"`
	Summary     string         `json:"summary"`
	CoverImage  string         `json:"cover_image"`
	AuthorID    uint           `gorm:"not null" json:"author_id"`
	Author      User           `gorm:"foreignKey:AuthorID" json:"author"`
	CategoryID  uint           `json:"category_id"`
	Category    *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags        []Tag          `gorm:"many2many:article_tags;" json:"tags,omitempty"`
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	IsPublished bool           `gorm:"default:false" json:"is_published"`
}

