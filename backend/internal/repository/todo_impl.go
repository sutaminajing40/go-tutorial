package repository

import (
	"go-tutorial/backend/internal/domain"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) FindAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) FindByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Todo{}, id).Error
}
