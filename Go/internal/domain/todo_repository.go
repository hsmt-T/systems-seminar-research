package domain

type TodoRepository interface {
	Create(todo *Todo) error
	FindAll() ([]Todo, error)
	DeleteByID(id TodoID) (bool, error)
	UpdateByID(todo *Todo) (bool, error)
}
