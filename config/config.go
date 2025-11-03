package config

import (
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	OSS      OSSConfig
}

type OSSConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
	Domain          string // OSS域名，用于生成访问URL
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

type JWTConfig struct {
	Secret     string
	ExpireTime time.Duration
}

var AppConfig = Config{
	Server: ServerConfig{
		Port: ":8080",
	},
	Database: DatabaseConfig{
		Driver: "sqlite",
		DSN:    "blog.db",
	},
	JWT: JWTConfig{
		Secret:     "your-secret-key-change-in-production",
		ExpireTime: 24 * time.Hour,
	},
	OSS: OSSConfig{
		Endpoint:        getEnvOrDefault("OSS_ENDPOINT", ""),          // 从环境变量OSS_ENDPOINT读取，或在控制台"概览"中查看"外网访问Endpoint"
		AccessKeyID:     getEnvOrDefault("OSS_ACCESS_KEY_ID", ""),     // 从环境变量OSS_ACCESS_KEY_ID读取
		AccessKeySecret: getEnvOrDefault("OSS_ACCESS_KEY_SECRET", ""), // 从环境变量OSS_ACCESS_KEY_SECRET读取
		BucketName:      getEnvOrDefault("OSS_BUCKET_NAME", ""),       // 从环境变量OSS_BUCKET_NAME读取
		Domain:          getEnvOrDefault("OSS_DOMAIN", ""),            // 从环境变量OSS_DOMAIN读取，留空使用OSS默认域名
	},
}

// getEnvOrDefault 从环境变量获取值，如果不存在则返回默认值
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
