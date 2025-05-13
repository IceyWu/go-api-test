package main

import (
	"log"
	_ "test-api/docs" // 导入 swagger docs
	"test-api/internal/config"
	"test-api/internal/model"
	"test-api/internal/router"

	"github.com/gin-gonic/gin"
)

// @title           Test API
// @version         1.0
// @description     A test RESTful API server
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	// 初始化数据库
	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库表
	config.DB.AutoMigrate(&model.User{})

	r := gin.Default()

	// 初始化路由
	router.InitRouter(r)

	// 启动服务器
	r.Run(":8080")
}
