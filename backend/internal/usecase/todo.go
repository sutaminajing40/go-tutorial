package usecase

import "go-tutorial/backend/internal/domain"

// TodoUsecase ビジネスロジックのインターフェース
type TodoUsecase interface {
	CreateTodo(title string) (*domain.Todo, error)
	GetAllTodos() ([]domain.Todo, error)
	GetTodoByID(id uint) (*domain.Todo, error)
	UpdateTodo(id uint, title string, completed bool) (*domain.Todo, error)
	DeleteTodo(id uint) error
}
