package model

type Todo struct {
	ID     int
	Title  string
	Status string
}

// コンストラクタ
func NewTodo(title, status string) (*Todo, error) {
	task := &Todo{
		Title:  title,
		Status: status,
	}

	return task, nil
}
