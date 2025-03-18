package main

import (
	"database/sql"
	"log"
	"todo-list-api/internal/handlers"
	"todo-list-api/internal/repositories"
	"todo-list-api/internal/services"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
	todoRepository := repositories.NewSqliteTodoRepo(db)
	emailService := services.NewDefaultEmailService()
	todoHandler := handlers.NewTodoHandler(todoRepository, emailService)

	r.POST("/todos", todoHandler.CreateTodoHandler)
	r.GET("/todos", todoHandler.GetTodoHandler)
	r.GET("/todos/:id", todoHandler.GetTodoByIdHandler)
	r.PUT("/todos/:id", todoHandler.UpdateTodoHandler)
	r.DELETE("/todos/:id", todoHandler.DeleteTodoHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
