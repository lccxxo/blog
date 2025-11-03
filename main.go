package main

import (
	"log"

	"github.com/lccxxo/blog/config"
	"github.com/lccxxo/blog/database"
	"github.com/lccxxo/blog/routes"
)

func main() {
	// 初始化数据库
	database.InitDB()

	// 设置路由
	r := routes.SetupRouter()

	// 启动服务器
	log.Printf("服务器启动在端口 %s\n", config.AppConfig.Server.Port)
	if err := r.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
