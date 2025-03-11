package entity

type ToDo struct {
	ID          uint64 `json:"id"`
	Description string `json:"description" binding:"required"`
	Done        bool   `json:"done" default:"false"`
}
