package service

import (
	"test-api/internal/config"
	"test-api/internal/model"
)

// UserService 用户服务
type UserService struct{}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// GetUsers 获取所有用户
func (s *UserService) GetUsers(page, size int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	// 获取总数
	config.DB.Model(&model.User{}).Count(&total)

	// 分页查询
	offset := (page - 1) * size
	result := config.DB.Offset(offset).Limit(size).Find(&users)
	
	return users, total, result.Error
}

// GetUser 获取单个用户
func (s *UserService) GetUser(id uint) (model.User, error) {
	var user model.User
	result := config.DB.First(&user, id)
	return user, result.Error
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(user *model.User) error {
	return config.DB.Save(user).Error
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}
