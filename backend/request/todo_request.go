package request

import "time"

type CreateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryID  int64      `json:"category_id"`
	Priority    string     `json:"priority"`
	Completed   bool       `json:"completed"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryID  int64      `json:"category_id"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	Completed   bool       `json:"completed"`
}
