package infra

import (
	"hello-gin/domain/model"
	"hello-gin/domain/repository"

	"github.com/jinzhu/gorm"
)

type TodoRepository struct {
	Conn *gorm.DB
}

// コンストラクタ
func NewTodoRepository(conn *gorm.DB) repository.TodoRepository {
	return &TodoRepository{Conn: conn}
}

// Create
func (tr *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	if err := tr.Conn.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}
