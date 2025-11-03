package config

import (
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
		Endpoint:        "oss-cn-beijing.aliyuncs.com", // OSS Endpoint，在控制台"概览"中查看"外网访问Endpoint"
		AccessKeyID:     "LTAI5tGUivbGfUY94gkBQ4K1",
		AccessKeySecret: "2bp4xRlpfxllU8SH8ZmLeADM3lW4RW",
		BucketName:      "zcc-mjh",
		Domain:          "", // 留空使用OSS默认域名，或填写自定义域名/CDN域名（如：https://cdn.example.com）
	},
}
