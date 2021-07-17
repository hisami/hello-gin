package usecase

import (
	"hello-gin/domain/model"
	"hello-gin/domain/repository"
)

type TodoUsecase interface {
	Create(title, status string) (*model.Todo, error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

// コンストラクタ
func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepo: todoRepo}
}

// Create
func (tu *todoUsecase) Create(title, status string) (*model.Todo, error) {
	todo, err := model.NewTodo(title, status)
	if err != nil {
		return nil, err
	}

	createdTodo, err := tu.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return createdTodo, nil
}
