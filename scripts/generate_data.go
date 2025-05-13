package main

import (
	"fmt"
	"math/rand"
	"time"

	"test-api/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `gorm:"size:50;not null"`
	Email     string `gorm:"size:100;not null;unique"`
	Age       int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		panic("Failed to load config: " + err.Error())
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(config.GlobalConfig.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 自动迁移
	db.AutoMigrate(&User{})

	// 生成随机数据
	rand.Seed(time.Now().UnixNano())
	users := make([]User, 1000)

	for i := 0; i < 1000; i++ {
		users[i] = User{
			Username: fmt.Sprintf("user%d", i+1),
			Email:    fmt.Sprintf("user%d@example.com", i+1),
			Age:      rand.Intn(50) + 18, // 18-67岁
		}
	}

	// 批量插入数据
	result := db.CreateInBatches(users, 100) // 每批100条
	if result.Error != nil {
		panic("Failed to insert data")
	}

	fmt.Printf("Successfully inserted %d records\n", result.RowsAffected)
}
