package persistence

import (
	"database/sql"

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

func (r *todoRepository) DeleteByID(id domain.TodoID) (bool, error) {
	query := `
		DELETE FROM todos
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *todoRepository) UpdateByID(todo *domain.Todo) (bool, error) {
	query := `
		UPDATE todos
		SET title = $2, description = $3
		WHERE id = $1
		RETURNING created_at
	`

	err := r.db.QueryRow(
		query,
		todo.ID,
		todo.Title,
		todo.Description,
	).Scan(&todo.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
