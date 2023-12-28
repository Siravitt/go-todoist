package todo_service

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   int    `json:"completed"`
}

type TodoService interface {
	GetTodos(int) ([]TodoResponse, error)
	GetTodo(int) (*TodoResponse, error)
	AddTodo(TodoRequest) (*TodoResponse, error)
	// UpdateTodo(TodoRequest) error
	// DeleteTodo(int)
}
