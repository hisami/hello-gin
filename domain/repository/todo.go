package repository

import (
	"hello-gin/domain/model"
)

type TodoRepository interface {
	Create(todo *model.Todo) (*model.Todo, error)
}
