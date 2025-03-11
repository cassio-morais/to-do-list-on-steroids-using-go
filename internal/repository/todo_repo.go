package repository

import (
	"database/sql"
	"fmt"
	"todo-list-api/internal/core/entity"

	_ "modernc.org/sqlite"
)

type todoRepo struct {
	DB *sql.DB
}

func NewTodoRepo(db *sql.DB) *todoRepo {
	return &todoRepo{
		DB: db,
	}
}

func (t *todoRepo) CreateTodo(todo *entity.ToDo) (int, error) {

	stmt, err := t.DB.Prepare("INSERT INTO todos (description, done) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var result sql.Result

	// the stmt.Exec receive values based on the VALUES order in the statement
	// description first, done second
	// smt.Exec is for database modifications, commands.
	result, err = stmt.Exec(&todo.Description, &todo.Done)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	if rows == 0 {
		return 0, fmt.Errorf("erro ao criar todo")
	}

	fmt.Println("Todo criado com sucesso!")

	return int(rows), nil
}

func (t *todoRepo) GetTodo(offset int, limit int) ([]entity.ToDo, error) {

	if offset <= 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = 10
	}

	var entities []entity.ToDo

	// create the query statement
	stmt, err := t.DB.Prepare("SELECT * FROM todos LIMIT ? OFFSET ?")
	if err != nil {
		return entities, err
	}
	defer stmt.Close()

	// append data to query statement and execute it
	// stmt.Query is for queries
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return entities, err
	}

	defer rows.Close() // close the rows (preventing lock)

	for rows.Next() {

		entity := entity.ToDo{}
		err = rows.Scan(&entity.ID, &entity.Description, &entity.Done)
		if err != nil {
			return entities, err
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (t *todoRepo) GetTodoById(id int) (entity.ToDo, error) {
	var entity entity.ToDo

	stmt, err := t.DB.Prepare("SELECT * FROM todos WHERE id = ?")
	if err != nil {
		return entity, err
	}
	defer stmt.Close() // close the statement

	rows, err := stmt.Query(id)

	if err != nil {
		return entity, err
	}

	defer rows.Close() // close the rows (preventing lock)

	if !rows.Next() {
		return entity, fmt.Errorf("todo não encontrado")
	}

	err = rows.Scan(&entity.ID, &entity.Description, &entity.Done)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (t *todoRepo) UpdateTodoById(id int, todo *entity.ToDo) (int, error) {
	stmt, err := t.DB.Prepare("UPDATE todos SET description = ?, done = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("todo não encontrado para atualização")
	}

	fmt.Println("Todo atualizado com sucesso!")

	return int(rowsAffected), nil
}

func (t *todoRepo) DeleteTodoById(id int) (int, error) {
	stmt, err := t.DB.Prepare("DELETE FROM todos WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("todo não encontrado para deleção")
	}

	fmt.Println("Todo deletado com sucesso!")

	return int(rowsAffected), nil
}
