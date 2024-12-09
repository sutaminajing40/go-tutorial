package repository

import "backend/internal/domain"

// TodoRepository データベース操作のインターフェース
type TodoRepository interface {
	Create(todo *domain.Todo) error
	FindAll() ([]domain.Todo, error)
	FindByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
}
