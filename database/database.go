package database

import (
	"log"

	"github.com/lccxxo/blog/config"
	"github.com/lccxxo/blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.AppConfig.Database.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Category{},
		&models.Tag{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database initialized successfully")
}
