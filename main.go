package main

import (
	"log"

	"github.com/lccxxo/blog/config"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/routes"
	"github.com/lccxxo/blog/utils"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 初始化OSS（如果配置了OSS）
	if err := utils.InitOSS(); err != nil {
		log.Printf("警告: OSS初始化失败: %v (将使用本地存储)\n", err)
	} else {
		log.Println("OSS初始化成功")
	}

	// 设置路由
	r := routes.SetupRouter()

	// 启动服务器
	log.Printf("服务器启动在端口 %s\n", config.AppConfig.Server.Port)
	if err := r.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
