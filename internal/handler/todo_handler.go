package handler

import (
	"strconv"
	"todo-list-api/internal/core/contract"
	"todo-list-api/internal/core/entity"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	toDoRepository contract.ToDoRepository
}

func NewTodoHandler(toDoRepository contract.ToDoRepository) *TodoHandler {
	return &TodoHandler{
		toDoRepository: toDoRepository,
	}
}

func (th *TodoHandler) CreateTodoHandler(ctx *gin.Context) {
	var todo entity.ToDo

	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := th.toDoRepository.CreateTodo(&todo)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, CreatedTodoResponse{Description: todo.Description, Done: todo.Done})
}

func (th *TodoHandler) GetTodoHandler(ctx *gin.Context) {
	offsetStr := ctx.Query("offset")
	limitStr := ctx.Query("limit")

	offset, _ := strconv.Atoi(offsetStr)
	limit, _ := strconv.Atoi(limitStr)

	todos, err := th.toDoRepository.GetTodo(offset, limit)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, todos)
}

func (th *TodoHandler) GetTodoByIdHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	todo, err := th.toDoRepository.GetTodoById(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, todo)
}

func (th *TodoHandler) UpdateTodoHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var todo entity.ToDo
	err := ctx.BindJSON(&todo)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = th.toDoRepository.UpdateTodoById(id, &todo)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	th.GetTodoByIdHandler(ctx)
}

func (th *TodoHandler) DeleteTodoHandler(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := th.toDoRepository.DeleteTodoById(id)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(204)
}
