package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"todo-list-api/internal/core/entities"

	_ "modernc.org/sqlite"
)

type SqliteTodoRepo struct {
	DB *sql.DB
}

func NewSqliteTodoRepo(db *sql.DB) *SqliteTodoRepo {
	return &SqliteTodoRepo{
		DB: db,
	}
}

func (t *SqliteTodoRepo) CreateTodo(todo *entities.ToDo) error {

	stmt, err := t.DB.Prepare("INSERT INTO todos (description, done) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// the stmt.Exec receive values based on the VALUES order in the statement
	// description first, done second
	// smt.Exec is for database modifications, commands.
	result, err := stmt.Exec(&todo.Description, &todo.Done)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("error when creating todo")
	}

	log.Printf("todo: %+v created", todo) // %+v and %v print the values of struct

	return nil
}

func (t *SqliteTodoRepo) GetTodo(offset int, limit int) ([]entities.ToDo, error) {

	if offset <= 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = 10
	}

	var todos []entities.ToDo

	// create the query statement
	stmt, err := t.DB.Prepare("SELECT * FROM todos LIMIT ? OFFSET ?")
	if err != nil {
		return todos, err
	}
	defer stmt.Close()

	// append data to query statement and execute it
	// stmt.Query is for queries
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return todos, err
	}

	defer rows.Close() // close the rows (preventing lock)

	for rows.Next() {

		todo := entities.ToDo{}
		err = rows.Scan(&todo.ID, &todo.Description, &todo.Done)
		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}

	log.Printf("todos: %+v", todos)

	return todos, nil
}

func (t *SqliteTodoRepo) GetTodoById(id int) (entities.ToDo, error) {
	var todo entities.ToDo

	stmt, err := t.DB.Prepare("SELECT * FROM todos WHERE id = ?")
	if err != nil {
		return todo, err
	}
	defer stmt.Close() // close the statement

	rows, err := stmt.Query(id)

	if err != nil {
		return todo, err
	}

	defer rows.Close() // close the rows (preventing lock)

	if !rows.Next() {
		return todo, fmt.Errorf("todo not found")
	}

	err = rows.Scan(&todo.ID, &todo.Description, &todo.Done)
	if err != nil {
		return todo, err
	}

	log.Printf("todo: %+v ", todo)

	return todo, nil
}

func (t *SqliteTodoRepo) UpdateTodoById(id int, todo *entities.ToDo) error {
	stmt, err := t.DB.Prepare("UPDATE todos SET description = ?, done = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Description, todo.Done, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}

	todo.ID = uint64(id)
	log.Printf("todo updated: %+v", todo)

	return nil
}

func (t *SqliteTodoRepo) DeleteTodoById(id int) error {
	stmt, err := t.DB.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}

	log.Printf("Todo deleted id: %+v", id)

	return nil
}
