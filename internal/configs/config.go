package configs

import (
	"log"

	"github.com/thanhpv3380/api/env"
)

type Redis struct {
	Host     string
	Port     int
	Password string
}

type Config struct {
	Port  int
	Redis Redis
}

var Cfg *Config

func LoadConfig() *Config {
	if err := env.LoadEnv(); err != nil {
		log.Println("No .env file found")
	}

	Cfg = &Config{
		Port: env.GetInt("PORT", 3000),
		Redis: Redis{
			Host:     env.GetString("REDIS_HOST", ""),
			Port:     env.GetInt("REDIS_PORT", 6379),
			Password: env.GetString("REDIS_PASSWORD", ""),
		},
	}

	log.Println("Load config successfully")

	return Cfg
}
