package app

import (
	"os"

	grpcapp "github.com/GosMachine/ProductService/internal/app/grpc"
	product "github.com/GosMachine/ProductService/internal/services"
	"github.com/GosMachine/ProductService/internal/storage/database"
	"github.com/GosMachine/ProductService/internal/storage/redis"
	"go.uber.org/zap"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *zap.Logger) *App {
	db, err := database.New()
	if err != nil {
		panic(err)
	}
	redis := redis.New(db)
	authService := product.New(log, db, redis)
	grpcApp := grpcapp.New(log, authService, os.Getenv("PRODUCT_SERVICE_ADDR"))
	return &App{
		GRPCSrv: grpcApp,
	}
}
