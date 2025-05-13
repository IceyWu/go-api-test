package router

import (
	"test-api/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	// API 版本分组
	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		userGroup := v1.Group("/users")
		{
			userGroup.GET("", handler.GetUsers)
			userGroup.GET("/:id", handler.GetUser)
			userGroup.POST("", handler.CreateUser)
			userGroup.PUT("/:id", handler.UpdateUser)
			userGroup.DELETE("/:id", handler.DeleteUser)
		}
	}

	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
} 
