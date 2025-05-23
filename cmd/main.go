package main

import (
	"fmt"

	"github.com/thanhpv3380/execution-producer/internal/configs"
	"github.com/thanhpv3380/execution-producer/internal/infra/redis"
	"github.com/thanhpv3380/execution-producer/internal/routers"

	logger "github.com/thanhpv3380/go-common/logger"
	server "github.com/thanhpv3380/go-common/server"
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
