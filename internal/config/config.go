package config

import (
	"fmt"
	"net"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// LogConfigSearcher interface for search Log config.
type LogConfigSearcher interface {
	Get() (*LogConfig, error)
}

// PgConfigSearcher interface for search PG config.
type PgConfigSearcher interface {
	Get() (*PgConfig, error)
}

// RedisConfigSearcher interface for search Redis config.
type RedisConfigSearcher interface {
	Get() (*RedisConfig, error)
}

// HTTPConfigSearcher interface for search HTTP config.
type HTTPConfigSearcher interface {
	Get() (*HTTPConfig, error)
}

// SwaggerConfigSearcher interface for search Swagger
type SwaggerConfigSearcher interface {
	Get() (*SwaggerConfig, error)
}

// Load dotenv from path to env
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

// LogConfig config for zerolog.
type LogConfig struct {
	LogLevel   zerolog.Level
	TimeFormat string
}

// PgConfig config for postgresql.
type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// RedisConfig config for Redis.
type RedisConfig struct {
	Host        string
	Port        string
	MaxIdle     int
	IdleTimeout time.Duration
	TTL         time.Duration
}

// Address ...
func (cfg *RedisConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

// DSN ...
func (cfg *PgConfig) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)
}

// HTTPConfig is config for HTTP
type HTTPConfig struct {
	Host string
	Port string
}

// Address get address from config
func (cfg *HTTPConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

// SwaggerConfig is config for Swagger
type SwaggerConfig struct {
	Host string
	Port string
}

// Address get address from config
func (cfg *SwaggerConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}
