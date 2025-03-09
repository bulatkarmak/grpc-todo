package main

import (
	"database/sql"
	"fmt"
	pb "github.com/bulatkarmak/grpc-todo/api/todo-list"
	"github.com/bulatkarmak/grpc-todo/internal/config"
	"github.com/bulatkarmak/grpc-todo/internal/handlers"
	"github.com/bulatkarmak/grpc-todo/internal/repository"
	"github.com/bulatkarmak/grpc-todo/internal/service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config.LoadConfig("values_local.yaml")

	db, err := sql.Open("postgres",
		fmt.Sprintf("user=%v dbname=%v sslmode=%v",
			config.AppConfig.Database.User,
			config.AppConfig.Database.DBName,
			config.AppConfig.Database.SSLMode))
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	listener, err := net.Listen(config.AppConfig.Server.Protocol, config.AppConfig.Server.Port)
	if err != nil {
		log.Fatalf("Ошибка прослушивания порта: %v", err)
	}

	repo := repository.NewToDoRepository(db)
	serv := service.NewToDoService(repo)
	handler := handlers.NewToDoHandler(serv)

	// gRPC-сервер будет обрабатывать входящие на localhost:50051 запросы
	grpcServer := grpc.NewServer()

	pb.RegisterToDoServiceServer(grpcServer, handler)

	fmt.Println("Сервер слушает на порту :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка работы сервера: %v", err)
	}
}
