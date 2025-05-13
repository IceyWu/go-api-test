package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"size:50;not null" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Age       int       `gorm:"not null" json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse 用户响应结构
type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
