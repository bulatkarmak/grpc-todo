package main

import (
	"database/sql"
	"fmt"
	pb "github.com/bulatkarmak/grpc-todo/api/todo-list"
	"github.com/bulatkarmak/grpc-todo/internal/handlers"
	"github.com/bulatkarmak/grpc-todo/internal/repository"
	"github.com/bulatkarmak/grpc-todo/internal/service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db, err := sql.Open("postgres", "user=bjustice dbname=grpc_todo sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	listener, err := net.Listen("tcp", ":50051")
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
