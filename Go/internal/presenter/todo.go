package presenter

import (
	"systems-seminar-research-go/internal/domain"
	"systems-seminar-research-go/internal/service"
)

type CreateTodoPresenter struct{}

func NewCreateTodoPresenter() *CreateTodoPresenter {
	return &CreateTodoPresenter{}
}

func (p *CreateTodoPresenter) Output(todo domain.Todo) service.CreateTodoOutput {
	return service.CreateTodoOutput{
		ID:          string(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}
}