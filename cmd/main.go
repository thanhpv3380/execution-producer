package main

import (
	"execution-producer/internal/configs"
	"execution-producer/internal/routers"
	"fmt"

	"github.com/thanhpv3380/api/infra/redis"

	logger "github.com/thanhpv3380/api/logger"
	server "github.com/thanhpv3380/api/server"
)

func main() {
	cfg := configs.LoadConfig()
	logger.NewLogger(nil)

	initRedis(cfg)

	app := server.NewServer(cfg.Port)

	routers.SetupRoutes(app)

	logger.Infof("🚀 Server running at port: %d", cfg.Port)
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}

func initRedis(cfg *configs.Config) {
	redisAddress := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	logger.Infof("Initing Redis client ..., address: %s", redisAddress)
	err := redis.NewClient(redisAddress, cfg.Redis.Password)
	if err != nil {
		logger.Fatal("Failed to connect redis", err)
	}

	logger.Info("Init Redis client successfully")
}
