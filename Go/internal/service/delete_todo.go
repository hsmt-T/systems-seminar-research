package service

import (
	"errors"

	"systems-seminar-research-go/internal/domain"
)

var ErrTodoNotFound = errors.New("todo not found")

type DeleteTodoInput struct {
	ID string
}

type DeleteTodoUseCase interface {
	Execute(input DeleteTodoInput) error
}

type deleteTodoInteractor struct {
	TodoRepo domain.TodoRepository
}

func NewDeleteTodoInteractor(todoRepo domain.TodoRepository) DeleteTodoUseCase {
	return &deleteTodoInteractor{
		TodoRepo: todoRepo,
	}
}

func (i *deleteTodoInteractor) Execute(input DeleteTodoInput) error {
	if input.ID == "" {
		return errors.New("id is required")
	}

	deleted, err := i.TodoRepo.DeleteByID(domain.TodoID(input.ID))
	if err != nil {
		return err
	}
	if !deleted {
		return ErrTodoNotFound
	}

	return nil
}
