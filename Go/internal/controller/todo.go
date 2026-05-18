package controller

import (
	"net/http"

	"systems-seminar-research-go/internal/service"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	createTodoUseCase  service.CreateTodoUseCase
	findAllTodoUseCase service.FindAllTodoUseCase
}

func NewTodoController(
	createTodoUseCase service.CreateTodoUseCase,
	findAllTodoUseCase service.FindAllTodoUseCase,
) *TodoController {
	return &TodoController{
		createTodoUseCase:  createTodoUseCase,
		findAllTodoUseCase: findAllTodoUseCase,
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

func (c *TodoController) FindAllTodoHandler(ctx echo.Context) error {
	output, err := c.findAllTodoUseCase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, output)
}
