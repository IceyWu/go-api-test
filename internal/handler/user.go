package handler

import (
	"net/http"
	"strconv"
	"test-api/internal/model"
	"test-api/internal/service"

	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

// GetUsers godoc
// @Summary 获取所有用户
// @Description 获取所有用户列表
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "页码，默认1" default(1)
// @Param size query int false "每页数量，默认10" default(10)
// @Success 200 {object} model.UserResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	// 获取分页参数
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		size = 10
	}

	users, total, err := userService.GetUsers(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Code:    500,
			Message: "Failed to get users",
		})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data: gin.H{
			"list":  users,
			"total": total,
			"page":  page,
			"size":  size,
		},
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	user, err := userService.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.UserResponse{
			Code:    404,
			Message: "User not found",
		})
		return
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
			Message: "Invalid request body",
		})
		return
	}

	if err := userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Code:    500,
			Message: "Failed to create user",
		})
		return
	}

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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Code:    400,
			Message: "Invalid request body",
		})
		return
	}

	user.ID = uint(id)
	if err := userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Code:    500,
			Message: "Failed to update user",
		})
		return
	}

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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Code:    400,
			Message: "Invalid user ID",
		})
		return
	}

	if err := userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Code:    500,
			Message: "Failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{
		Code:    200,
		Message: "success",
		Data:    "User deleted successfully",
	})
}
