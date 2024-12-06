package handler

import "github.com/labstack/echo/v4"

// TodoHandler HTTPハンドラーのインターフェース
type TodoHandler interface {
	CreateTodo() echo.HandlerFunc
	GetAllTodos() echo.HandlerFunc
	GetTodoByID() echo.HandlerFunc
	UpdateTodo() echo.HandlerFunc
	DeleteTodo() echo.HandlerFunc
}

// CreateTodoRequest Todoの作成リクエスト
type CreateTodoRequest struct {
	Title string `json:"title"`
}

// UpdateTodoRequest Todoの更新リクエスト
type UpdateTodoRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
