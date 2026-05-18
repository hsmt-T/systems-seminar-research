package service

import (
	"errors"
	"systems-seminar-research-go/internal/domain"
	"time"
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

type CreateTodoPresenter interface {
	Output(domain.Todo) CreateTodoOutput
}

type createTodoInteractor struct {
	TodoRepo      domain.TodoRepository
	TodoPresenter CreateTodoPresenter
}

func NewCreateTodoInteractor(todoRepo domain.TodoRepository, todoPresenter CreateTodoPresenter) CreateTodoUseCase {
	return &createTodoInteractor{
		TodoRepo:      todoRepo,
		TodoPresenter: todoPresenter,
	}
}

func (i *createTodoInteractor) Execute(input CreateTodoInput) (CreateTodoOutput, error) {
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

	return i.TodoPresenter.Output(todo), nil
}
