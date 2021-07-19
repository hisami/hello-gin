package model

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// コンストラクタ
func NewTodo(title, status string) (*Todo, error) {
	task := &Todo{
		Title:  title,
		Status: status,
	}

	return task, nil
}
