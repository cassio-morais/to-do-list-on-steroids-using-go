package contracts

import "todo-list-api/internal/core/entities"

type ToDoRepository interface {
	CreateTodo(todo *entities.ToDo) error
	GetTodo(offset int, limit int) (todos []entities.ToDo, err error)
	GetTodoById(id int) (todo entities.ToDo, err error)
	UpdateTodoById(id int, todo *entities.ToDo) error
	DeleteTodoById(id int) error
}
