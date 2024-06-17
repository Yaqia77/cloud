package mysql

import (
	"cloud/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     int    `toml:"port" env:"MYSQL_PORT"`
	Username string `toml:"user" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`

	//MaxOpenConns是连接池中最大连接数，默认值为0表示不限制
	MaxOpenConns int `toml:"max_open_conns" env:"MYSQL_MAX_OPEN_CONNS"`
	//MaxIdleConns是连接池中最大空闲连接数，默认值为2
	//当开启空闲连接回收时，该值才会生效，并发量大时建议设置该值，以防止连接过多导致连接池溢出
	MaxIdleConns int `toml:"max_idle_conns" env:"MYSQL_MAX_IDLE_CONNS"`
}

var (
	DB *gorm.DB
)

type Option func(*MySQL)

func MaxOpenConns(maxOpenConns int) Option {
	return func(o *MySQL) {
		o.MaxOpenConns = maxOpenConns
	}
}

func MaxIdleConns(maxIdleConns int) Option {
	return func(o *MySQL) {
		o.MaxIdleConns = maxIdleConns
	}
}

// NewMySQL creates a new MySQL instance with the given options.
// 配置加载
// 读取环境变量,初始化数据库连接,设置最大连接数和最大空闲连接数,设置日志级别,返回MySQL实例
func NewMySQL(opts ...Option) *MySQL {
	mysql := &MySQL{
		Host:         getEnv("MYSQL_HOST", "localhost"),
		Port:         getEnvAsInt("MYSQL_PORT", 3306),
		Username:     getEnv("MYSQL_USERNAME", "username"),
		Password:     getEnv("MYSQL_PASSWORD", "123456"),
		Database:     getEnv("MYSQL_DATABASE", "cloud_storage"),
		MaxOpenConns: getEnvAsInt("MYSQL_MAX_OPEN_CONNS", 10),
		MaxIdleConns: getEnvAsInt("MYSQL_MAX_IDLE_CONNS", 5),
	}
	for _, opt := range opts {
		opt(mysql)
	}
	return mysql
}

func (m *MySQL) InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移表结构
	models.AutoMigrateUserInfoTable(DB)
	models.AutoMigrateFileInfoTable(DB)
}

// 读取环境变量的实用程序函数
func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// 读取环境变量并转换为int的实用程序函数
func getEnvAsInt(key string, defaultValue int) int {
	if valueStr, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}
