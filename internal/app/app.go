package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/ander320/authorization-service/internal/app/grpc"
	"github.com/ander320/authorization-service/internal/services/auth"
	"github.com/ander320/authorization-service/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	//initialize storage
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	//initialize auth service
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	//initialize grpc app
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
