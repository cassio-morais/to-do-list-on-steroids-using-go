package contract

import "todo-list-api/internal/core/entity"

type ToDoRepository interface {
	CreateTodo(todo *entity.ToDo) error
	GetTodo(offset int, limit int) (todos []entity.ToDo, err error)
	GetTodoById(id int) (todo entity.ToDo, err error)
	UpdateTodoById(id int, todo *entity.ToDo) error
	DeleteTodoById(id int) error
}
