package services

import (
	"context"
	"math/rand"
	"time"

	"gitlab.com/trivery-id/skadi/internal/todo/domain"
)

type TodoService struct {
	todos []domain.Todo
}

var Service = &TodoService{
	todos: []domain.Todo{},
}

func (svc *TodoService) CreateNewTodo(ctx context.Context, in CreateNewTodoInput) (*domain.Todo, error) {
	todo := domain.Todo{
		ID:          uint64(rand.Int()),
		CreatedAt:   time.Now(),
		Text:        in.Text,
		UserID:      in.UserID,
		Description: in.Description,
	}

	svc.todos = append(svc.todos, todo)
	return &todo, nil
}

func (svc *TodoService) GetAllTodos(ctx context.Context, userID uint64) ([]domain.Todo, error) {
	out := []domain.Todo{}
	for _, todo := range svc.todos {
		if todo.UserID == userID {
			out = append(out, todo)
		}
	}

	return out, nil
}