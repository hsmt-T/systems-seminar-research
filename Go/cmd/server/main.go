package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"systems-seminar-research-go/internal/controller"
	"systems-seminar-research-go/internal/infrastructure/database"
	"systems-seminar-research-go/internal/infrastructure/persistence"
	"systems-seminar-research-go/internal/service"
)

type appState struct {
	db *sqlx.DB
}

func main() {
	databaseURL := getenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/app?sslmode=disable")

	dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := database.OpenPostgres(dbCtx, databaseURL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	// Initialize dependencies
	todoRepo := persistence.NewTodoRepository(db)
	createTodoUseCase := service.NewCreateTodoInteractor(todoRepo)
	findAllTodoUseCase := service.NewFindAllTodoInteractor(todoRepo)
	deleteTodoUseCase := service.NewDeleteTodoInteractor(todoRepo)
	todoController := controller.NewTodoController(createTodoUseCase, findAllTodoUseCase, deleteTodoUseCase)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", rootHandler)
	e.GET("/todos", todoController.FindAllTodoHandler)
	e.POST("/todos", todoController.CreateTodoHandler)
	e.DELETE("/todos/:id", todoController.DeleteTodoHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Go Echo + sqlx + PostgreSQL")
}
