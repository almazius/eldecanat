package main

import (
	"eldick/eldick/config"
	"eldick/eldick/internal/auth/http"
	"eldick/eldick/internal/auth/repository/postgtres"
	redis2 "eldick/eldick/internal/auth/repository/redis"
	"eldick/eldick/internal/auth/service"
	http2 "eldick/eldick/internal/service/http"
	"eldick/eldick/pkg/logger"
	"eldick/eldick/pkg/postrges"
	"eldick/eldick/pkg/redis"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func main() {
	logger.Init()

	v, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed load config", "error", err)
		return
	}

	err = config.ParseConfig(v)
	if err != nil {
		slog.Error("failed parse config", "error", err)
		return
	}

	Init()
}

func Init() {
	db, err := postrges.InitPsqlDB(config.C())
	if err != nil {
		slog.Error("failed init postgres db", "error", err)
		panic(err)
	}

	rdb, err := redis.Init(config.C())
	if err != nil {
		slog.Error("failed init redis db", "error", err)
		panic(err)
	}

	authCache := redis2.NewAuthCache(rdb)
	authRepository := postgtres.NewAuthRepository(db)

	authService := service.NewAuthService(authRepository, authCache)

	app := fiber.New()

	authHandler := http.NewAuthHandler(authService)
	mw := http.NewAuthMW(authService)

	authGroup := app.Group("/auth")
	apiGroup := app.Group("/api")

	http.AuthRoute(authGroup, authHandler)
	http2.ServiceRoute(apiGroup, mw, nil)
	err = app.Listen(fmt.Sprintf(":%d", config.C().Port))
	if err != nil {
		slog.Error("failed listen port for start service", "error", err,
			slog.Int("port", config.C().Port))
		return
	}
}
