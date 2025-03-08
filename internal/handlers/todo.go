package handlers

import (
	"context"
	pb "github.com/bulatkarmak/grpc-todo/api/todo-list"
	"github.com/bulatkarmak/grpc-todo/internal/domain"
	"github.com/bulatkarmak/grpc-todo/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type ToDoHandler struct {
	pb.UnimplementedToDoServiceServer
	service service.ToDoService
}

func NewToDoHandler(service service.ToDoService) *ToDoHandler {
	return &ToDoHandler{service: service}
}

func (h *ToDoHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	if req.Title == "" {
		log.Printf("ошибка: title не может быть пустым")
		return nil, status.Error(codes.InvalidArgument, "title не может быть пустым")
	}

	if req.Description == "" {
		log.Printf("ошибка: description не может быть пустым")
		return nil, status.Error(codes.InvalidArgument, "description не может быть пустым")
	}

	params := &domain.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
	}

	task, err := h.service.CreateTask(ctx, params)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка при создании task: %v", err)
	}

	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			TaskId:      task.ID,
			Title:       task.Title,
			Description: task.Description,
			IsCompleted: task.IsCompleted,
			CreatedAt:   timestamppb.New(task.CreatedAt),
			UpdatedAt:   timestamppb.New(task.UpdatedAt),
		},
	}, nil
}

func (h *ToDoHandler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	return nil, nil
}

func (h *ToDoHandler) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	return nil, nil
}

func (h *ToDoHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	return nil, nil
}

func (h *ToDoHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	return nil, nil
}
