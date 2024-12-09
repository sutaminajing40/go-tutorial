package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
	"strings"
)

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

// NewTodoUsecase ユースケースのインスタンスを生成
func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepo: todoRepo}
}

func (u *todoUsecase) CreateTodo(title string) (*domain.Todo, error) {
	if len(strings.TrimSpace(title)) == 0 || len(title) > 100 {
		return nil, errors.New("タイトルは1文字以上100文字以下である必要があります")
	}
	todo := &domain.Todo{Title: title}
	if err := u.todoRepo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) GetAllTodos() ([]domain.Todo, error) {
	return u.todoRepo.FindAll()
}

func (u *todoUsecase) GetTodoByID(id uint) (*domain.Todo, error) {
	if id == 0 {
		return nil, errors.New("IDは1以上である必要があります")
	}
	return u.todoRepo.FindByID(id)
}

func (u *todoUsecase) UpdateTodo(id uint, title string, completed bool) (*domain.Todo, error) {
	if id == 0 {
		return nil, errors.New("IDは1以上である必要があります")
	}
	if len(strings.TrimSpace(title)) == 0 || len(title) > 100 {
		return nil, errors.New("タイトルは1文字以上100文字以下である必要があります")
	}
	todo, err := u.todoRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	todo.Title = title
	todo.Completed = completed
	if err := u.todoRepo.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (u *todoUsecase) DeleteTodo(id uint) error {
	if id == 0 {
		return errors.New("IDは1以上である必要があります")
	}
	if _, err := u.todoRepo.FindByID(id); err != nil {
		return err
	}
	return u.todoRepo.Delete(id)
}
