package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 应用配置结构体
type Config struct {
	Database DatabaseConfig
}

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// GetDSN 获取数据库连接字符串
func (d *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Host, d.Port, d.DBName)
}

var GlobalConfig Config

// Init 初始化配置
func Init() error {
	// 设置配置文件路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 设置环境变量前缀
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 如果配置文件不存在，使用默认配置
			setDefaultConfig()
		} else {
			return err
		}
	}

	// 将配置映射到结构体
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		return err
	}

	return nil
}

// setDefaultConfig 设置默认配置
func setDefaultConfig() {
	viper.SetDefault("database.host", "127.0.0.1")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "123456")
	viper.SetDefault("database.dbname", "test_api")
}
