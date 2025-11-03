package config

import (
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
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
}

