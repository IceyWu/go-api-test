package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	// 初始化配置
	if err := Init(); err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// 首先连接到 MySQL 服务器（不指定数据库）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		GlobalConfig.Database.User,
		GlobalConfig.Database.Password,
		GlobalConfig.Database.Host,
		GlobalConfig.Database.Port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL server: %v", err)
	}

	// 创建数据库
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci",
		GlobalConfig.Database.DBName)).Error
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	// 关闭初始连接
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}
	sqlDB.Close()

	// 连接到新创建的数据库
	db, err = gorm.Open(mysql.Open(GlobalConfig.Database.GetDSN()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
	return nil
}
