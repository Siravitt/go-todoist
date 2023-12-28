package handler

import (
	"github.com/Siravitt/go-todoist/service/todo_service"
	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	todoSrv todo_service.TodoService
}

func NewTodoHandler(todoSrv todo_service.TodoService) todoHandler {
	return todoHandler{todoSrv: todoSrv}
}

func (h todoHandler) GetTodos(c echo.Context) error {
	return nil
}

func (h todoHandler) GetTodo(c echo.Context) error {
	return nil
}

func (h todoHandler) AddTodo(c echo.Context) error {
	return nil
}
