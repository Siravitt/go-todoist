package todo_service

import "github.com/Siravitt/go-todoist/repository/todo_repository"

type todoService struct {
	todoRepo todo_repository.TodoRepository
}

func NewTodoService(todoRepo todo_repository.TodoRepository) TodoService {
	return todoService{todoRepo: todoRepo}
}

// Get all todo from user_id
func (s todoService) GetTodos(userID int) ([]TodoResponse, error) {
	todos, err := s.todoRepo.GetAll(userID)
	if err != nil {
		return nil, err
	}

	todoResponses := []TodoResponse{}
	for _, todo := range todos {
		todoResponse := TodoResponse{
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		}
		todoResponses = append(todoResponses, todoResponse)
	}
	return todoResponses, nil
}

// Get todo by id
func (s todoService) GetTodo(id int) (*TodoResponse, error) {
	todo, err := s.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	todoResponse := TodoResponse{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	return &todoResponse, nil
}

// Create todo
func (s todoService) AddTodo(t TodoRequest) (*TodoResponse, error) {

	return nil, nil
}
