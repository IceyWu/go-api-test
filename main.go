package main

import (
	_ "test-api/docs" // 导入 swagger docs
	"test-api/internal/router"

	"github.com/gin-gonic/gin"
)

// @title           Test API
// @version         1.0
// @description     A test RESTful API server
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	r := gin.Default()

	// 初始化路由
	router.InitRouter(r)

	// 启动服务器
	r.Run(":8080")
}
