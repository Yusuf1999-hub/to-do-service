package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Yusuf1999-hub/to-do-service/config"
	pb "github.com/Yusuf1999-hub/to-do-service/genproto"
	"github.com/Yusuf1999-hub/to-do-service/pkg/db"
	"github.com/Yusuf1999-hub/to-do-service/pkg/logger"
	"github.com/Yusuf1999-hub/to-do-service/service"
	"github.com/Yusuf1999-hub/to-do/storage"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "to-do-services")
	defer func(l logger.Logger) {
		err := logger.Cleanup(l)
		if err != nil {
			log.Fatal("failed cleanup logger", logger.Error(err))
		}
	}(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	defer connDB.Close()

	pgStorage := storage.NewStoragePg(connDB)

	taskService := service.NewTaskService(pgStorage, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, taskService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
