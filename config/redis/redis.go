package redis

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
	Ctx    = context.Background() // 这是一个导出的变量
)

// Config stores Redis configuration
type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// InitRedis initializes the Redis client with the given configuration
func InitRedis(cfg *Config) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")
}

// LoadConfig loads the Redis configuration from environment variables
func LoadConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		port = 6379 // default port
	}
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		db = 0 // default DB
	}
	return &Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     port,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}
