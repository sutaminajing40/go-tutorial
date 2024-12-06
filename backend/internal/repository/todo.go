package repository

import "go-tutorial/backend/internal/domain"

// TodoRepository データベース操作のインターフェース
type TodoRepository interface {
	Create(todo *domain.Todo) error
	FindAll() ([]domain.Todo, error)
	FindByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
}
