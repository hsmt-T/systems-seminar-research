package controller

import (
	"net/http"

	"systems-seminar-research-go/internal/domain"
	"systems-seminar-research-go/internal/service"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	createTodoUseCase service.CreateTodoUseCase
}

type createTodoPresenter struct{}

func (p *createTodoPresenter) Output(todo domain.Todo) service.CreateTodoOutput {
	return service.CreateTodoOutput{
		ID:          string(todo.ID),
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}
}

func NewTodoController(createTodoUseCase service.CreateTodoUseCase) *TodoController {
	return &TodoController{
		createTodoUseCase: createTodoUseCase,
	}
}

func (c *TodoController) CreateTodoHandler(ctx echo.Context) error {
	var input service.CreateTodoInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	output, err := c.createTodoUseCase.Execute(input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, output)
}
