package usecase

import "go-tutorial/internal/domain"

// TodoUsecase ビジネスロジックのインターフェース
type TodoUsecase interface {
	CreateTodo(title string) error
	GetAllTodos() ([]domain.Todo, error)
	GetTodoByID(id uint) (*domain.Todo, error)
	UpdateTodo(id uint, title string, completed bool) error
	DeleteTodo(id uint) error
}
