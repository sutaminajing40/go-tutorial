package usecase

import (
	"errors"
	"go-tutorial/internal/domain"
	"go-tutorial/internal/repository"
	"strings"
)

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

// NewTodoUsecase ユースケースのインスタンスを生成
func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepo: todoRepo}
}

func (u *todoUsecase) CreateTodo(title string) error {
	if len(strings.TrimSpace(title)) == 0 || len(title) > 100 {
		return errors.New("タイトルは1文字以上100文字以下である必要があります")
	}
	todo := &domain.Todo{Title: title}
	return u.todoRepo.Create(todo)
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

func (u *todoUsecase) UpdateTodo(id uint, title string, completed bool) error {
	if id == 0 {
		return errors.New("IDは1以上である必要があります")
	}
	if len(strings.TrimSpace(title)) == 0 || len(title) > 100 {
		return errors.New("タイトルは1文字以上100文字以下である必要があります")
	}
	todo, err := u.todoRepo.FindByID(id)
	if err != nil {
		return err
	}
	todo.Title = title
	todo.Completed = completed
	return u.todoRepo.Update(todo)
}

func (u *todoUsecase) DeleteTodo(id uint) error {
	if id == 0 {
		return errors.New("IDは1以上である必要があります")
	}
	return u.todoRepo.Delete(id)
}
