package controller

import (
	"errors"
	"net/http"

	"systems-seminar-research-go/internal/service"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	createTodoUseCase  service.CreateTodoUseCase
	findAllTodoUseCase service.FindAllTodoUseCase
	deleteTodoUseCase  service.DeleteTodoUseCase
}

func NewTodoController(
	createTodoUseCase service.CreateTodoUseCase,
	findAllTodoUseCase service.FindAllTodoUseCase,
	deleteTodoUseCase service.DeleteTodoUseCase,
) *TodoController {
	return &TodoController{
		createTodoUseCase:  createTodoUseCase,
		findAllTodoUseCase: findAllTodoUseCase,
		deleteTodoUseCase:  deleteTodoUseCase,
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

func (c *TodoController) DeleteTodoHandler(ctx echo.Context) error {
	input := service.DeleteTodoInput{
		ID: ctx.Param("id"),
	}
	if input.ID == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "id is required"})
	}

	if err := c.deleteTodoUseCase.Execute(input); err != nil {
		if errors.Is(err, service.ErrTodoNotFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
