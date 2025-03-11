package contract

import "todo-list-api/internal/core/entity"

type ToDoRepository interface {
	CreateTodo(todo *entity.ToDo) (rowsAffected int, err error)
	GetTodo(offset int, limit int) (todos []entity.ToDo, err error)
	GetTodoById(id int) (todo entity.ToDo, err error)
	UpdateTodoById(id int, todo *entity.ToDo) (rowsAffected int, err error)
	DeleteTodoById(id int) (rowsAffected int, err error)
}
