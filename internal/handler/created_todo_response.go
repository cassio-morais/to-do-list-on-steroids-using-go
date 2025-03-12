package handler

type CreatedTodoResponse struct {
	Description string `json:"description" binding:"required"`
	Done        bool   `json:"done" default:"false"`
}
