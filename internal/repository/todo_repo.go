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
	// a passagem dos valores em .Exec segue a ordem dos campos na criação do statement
	// ou seja, o primeiro ? é o description e o segundo é o done
	// smt.Exec é para modificações no banco de dados, enquanto smt.Query é para consultas
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

	// cria a query
	stmt, err := t.DB.Prepare("SELECT * FROM todos LIMIT ? OFFSET ?")
	if err != nil {
		return entities, err
	}
	defer stmt.Close()

	// apenda os dados na query e executa
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return entities, err
	}

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
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return entity, err
	}

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

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows == 0 {
		return 0, fmt.Errorf("todo não encontrado para atualização")
	}

	fmt.Println("Todo atualizado com sucesso!")

	return int(rows), nil
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

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows == 0 {
		return 0, fmt.Errorf("todo não encontrado para deleção")
	}

	fmt.Println("Todo deletado com sucesso!")

	return int(rows), nil
}
