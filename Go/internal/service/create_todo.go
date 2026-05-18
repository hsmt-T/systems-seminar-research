package service

import (
	"errors"
	"time"

	"systems-seminar-research-go/internal/domain"
)

type CreateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTodoOutput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodoUseCase interface {
	Execute(input CreateTodoInput) (CreateTodoOutput, error)
}

type createTodoInteractor struct {
	TodoRepo domain.TodoRepository
}

func NewCreateTodoInteractor(
	todoRepo domain.TodoRepository,
) CreateTodoUseCase {
	return &createTodoInteractor{
		TodoRepo: todoRepo,
	}
}

func (i *createTodoInteractor) Execute(
	input CreateTodoInput,
) (CreateTodoOutput, error) {

	if input.Title == "" {
		return CreateTodoOutput{}, errors.New("title is required")
	}

	now := time.Now()

	todo := domain.NewTodo(
		domain.NewTodoId(),
		input.Title,
		input.Description,
		now,
	)

	if err := i.TodoRepo.Create(&todo); err != nil {
		return CreateTodoOutput{}, err
	}

	return CreateTodoOutput{
		ID:          string(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}, nil
}