package service

import (
	"time"

	"systems-seminar-research-go/internal/domain"
)

type FindAllTodoOutput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type FindAllTodoUseCase interface {
	Execute() ([]FindAllTodoOutput, error)
}

type findAllTodoInteractor struct {
	TodoRepo domain.TodoRepository
}

func NewFindAllTodoInteractor(todoRepo domain.TodoRepository) FindAllTodoUseCase {
	return &findAllTodoInteractor{
		TodoRepo: todoRepo,
	}
}

func (i *findAllTodoInteractor) Execute() ([]FindAllTodoOutput, error) {
	todos, err := i.TodoRepo.FindAll()
	if err != nil {
		return nil, err
	}

	output := make([]FindAllTodoOutput, 0, len(todos))
	for _, todo := range todos {
		output = append(output, FindAllTodoOutput{
			ID:          string(todo.ID),
			Title:       todo.Title,
			Description: todo.Description,
			CreatedAt:   todo.CreatedAt,
		})
	}

	return output, nil
}
