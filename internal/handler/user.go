package handler

import (
	"net/http"
	"test-api/internal/model"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary 获取所有用户
// @Description 获取所有用户列表
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.UserResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	// 模拟数据
	users := []model.User{
		{ID: "1", Username: "user1", Email: "user1@example.com", Age: 20},
		{ID: "2", Username: "user2", Email: "user2@example.com", Age: 25},
	}
	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data:    users,
	})
}

// GetUser godoc
// @Summary 获取单个用户
// @Description 根据ID获取用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// 模拟数据
	user := model.User{
		ID:       id,
		Username: "user" + id,
		Email:    "user" + id + "@example.com",
		Age:      20,
	}
	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data:    user,
	})
}

// CreateUser godoc
// @Summary 创建用户
// @Description 创建新用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "用户信息"
// @Success 200 {object} model.UserResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Code:    400,
			Message: "invalid request body",
		})
		return
	}
	// 模拟创建用户
	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data:    user,
	})
}

// UpdateUser godoc
// @Summary 更新用户
// @Description 更新用户信息
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Param user body model.User true "用户信息"
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Code:    400,
			Message: "invalid request body",
		})
		return
	}
	user.ID = id
	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data:    user,
	})
}

// DeleteUser godoc
// @Summary 删除用户
// @Description 删除指定用户
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data:    "User " + id + " deleted",
	})
}
