package service

import (
	"errors"
	"time"

	"systems-seminar-research-go/internal/domain"
)

type UpdateTodoInput struct {
	ID          string `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoOutput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateTodoUseCase interface {
	Execute(input UpdateTodoInput) (UpdateTodoOutput, error)
}

type updateTodoInteractor struct {
	TodoRepo domain.TodoRepository
}

func NewUpdateTodoInteractor(todoRepo domain.TodoRepository) UpdateTodoUseCase {
	return &updateTodoInteractor{
		TodoRepo: todoRepo,
	}
}

func (i *updateTodoInteractor) Execute(input UpdateTodoInput) (UpdateTodoOutput, error) {
	if input.ID == "" {
		return UpdateTodoOutput{}, errors.New("id is required")
	}
	if input.Title == "" {
		return UpdateTodoOutput{}, errors.New("title is required")
	}

	todo := domain.NewTodo(
		domain.TodoID(input.ID),
		input.Title,
		input.Description,
		time.Time{},
	)

	updated, err := i.TodoRepo.UpdateByID(&todo)
	if err != nil {
		return UpdateTodoOutput{}, err
	}
	if !updated {
		return UpdateTodoOutput{}, ErrTodoNotFound
	}

	return UpdateTodoOutput{
		ID:          string(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}, nil
}
