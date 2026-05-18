package persistence

import (
	"systems-seminar-research-go/internal/domain"

	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) domain.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {

	query := `
		INSERT INTO todos (id, title, description, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		query,
		todo.ID,
		todo.Title,
		todo.Description,
		todo.CreatedAt,
	)

	return err
}

func (r *todoRepository) FindAll() ([]domain.Todo, error) {
	todos := []domain.Todo{}

	query := `
		SELECT id, title, description, created_at
		FROM todos
		ORDER BY created_at DESC
	`

	if err := r.db.Select(&todos, query); err != nil {
		return nil, err
	}

	return todos, nil
}
