package main

import (
	"database/sql"
	"log"
	"todo-list-api/internal/handler"
	"todo-list-api/internal/repository"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela se não existir
	createTableSQL := `CREATE TABLE IF NOT EXISTS  todos (
    	id INTEGER PRIMARY KEY AUTOINCREMENT , 
    	description TEXT NOT NULL,
    	done BOOLEAN NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	r := gin.Default()
	r.Use(gin.Logger())

	// dependency injection
	todoRepository := repository.NewTodoRepo(db)
	handler := handler.NewTodoHandler(todoRepository)

	r.POST("/todos", handler.CreateTodoHandler)
	r.GET("/todos", handler.GetTodoHandler)
	r.GET("/todos/:id", handler.GetTodoByIdHandler)
	r.PUT("/todos/:id", handler.UpdateTodoHandler)
	r.DELETE("/todos/:id", handler.DeleteTodoHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
