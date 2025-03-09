package handlers

import (
	"context"
	"errors"
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
	params := &domain.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
	}

	task, err := h.service.CreateTask(ctx, params)

	if errors.Is(err, service.EmptyTitleErr) {
		log.Printf("ошибка: title не может быть пустым")
		return nil, status.Error(codes.InvalidArgument, "title не может быть пустым")
	}

	if errors.Is(err, service.EmptyDescErr) {
		log.Printf("ошибка: description не может быть пустым")
		return nil, status.Error(codes.InvalidArgument, "description не может быть пустым")
	}

	if err != nil {
		log.Printf("ошибка при создании task: %v", err)
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
	if req.TaskId == 0 {
		log.Printf("ошибка: taskID должен быть больше 0")
		return nil, status.Error(codes.InvalidArgument, "taskID должен быть больше 0")
	}

	taskID := req.TaskId

	task, err := h.service.GetTask(ctx, taskID)

	if err != nil {
		log.Printf("ошибка получения task: %v", err)
		return nil, status.Errorf(codes.Internal, "ошибка получения task: %v", err)
	}

	return &pb.GetTaskResponse{
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

func (h *ToDoHandler) ListTasks(ctx context.Context, _ *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, err := h.service.ListTasks(ctx)

	if err != nil {
		log.Printf("ошибка получения tasks: %v", err)
		return nil, status.Errorf(codes.Internal, "ошибка получения tasks: %v", err)
	}

	var pbTasks []*pb.Task
	for _, task := range tasks {
		pbTasks = append(pbTasks, &pb.Task{
			TaskId:      task.ID,
			Title:       task.Title,
			Description: task.Description,
			IsCompleted: task.IsCompleted,
			CreatedAt:   timestamppb.New(task.CreatedAt),
			UpdatedAt:   timestamppb.New(task.UpdatedAt),
		})
	}

	return &pb.ListTasksResponse{
		Tasks: pbTasks,
	}, nil
}

func (h *ToDoHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	params := &domain.UpdateTaskParams{
		ID:          req.TaskId,
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: req.IsCompleted,
	}

	task, err := h.service.UpdateTask(ctx, params)

	if errors.Is(err, service.EmptyTitleErr) {
		log.Printf("ошибка: title не может быть пустым")
		return nil, status.Error(codes.InvalidArgument, "title не может быть пустым")
	}

	if errors.Is(err, service.EmptyDescErr) {
		log.Printf("ошибка: description не может быть пустым")
		return nil, status.Error(codes.InvalidArgument, "description не может быть пустым")
	}

	if err != nil {
		log.Printf("ошибка при обновлении task: %v", err)
		return nil, status.Errorf(codes.Internal, "ошибка при обновлении task: %v", err)
	}

	return &pb.UpdateTaskResponse{
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

func (h *ToDoHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	return nil, nil
}
