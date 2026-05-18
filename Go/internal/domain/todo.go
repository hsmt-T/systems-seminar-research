package domain

import (
	"time"
)

type Todo struct {
	ID          TodoID    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

func NewTodo(id TodoID, title string, description string, createdAt time.Time) Todo {
	return Todo{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
	}
}
