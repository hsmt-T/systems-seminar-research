package domain

import (
	gouuid "github.com/satori/go.uuid"
)

type TodoID string

func NewTodoId() TodoID {
	return TodoID(gouuid.NewV4().String())
}