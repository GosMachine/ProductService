package app

import (
	grpcapp "github.com/GosMachine/ProductService/internal/app/grpc"
	"github.com/GosMachine/ProductService/internal/database/postgres"
	"github.com/GosMachine/ProductService/internal/database/redis"
	product "github.com/GosMachine/ProductService/internal/services"
	"go.uber.org/zap"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *zap.Logger, grpcPort int) *App {
	db, err := postgres.New()
	if err != nil {
		panic(err)
	}
	redis := redis.New(db)
	authService := product.New(log, db, redis)
	grpcApp := grpcapp.New(log, authService, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}
