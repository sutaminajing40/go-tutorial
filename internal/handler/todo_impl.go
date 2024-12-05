package handler

import (
	"net/http"
	"strconv"

	"go-tutorial/internal/usecase"

	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

// NewTodoHandler ハンドラーのインスタンスを生成
func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	return &todoHandler{todoUsecase: todoUsecase}
}

func (h *todoHandler) CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req CreateTodoRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なリクエストです"})
		}

		todo, err := h.todoUsecase.CreateTodo(req.Title)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, todo)
	}
}

func (h *todoHandler) GetAllTodos() echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := h.todoUsecase.GetAllTodos()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Todoの取得に失敗しました"})
		}

		return c.JSON(http.StatusOK, todos)
	}
}

func (h *todoHandler) GetTodoByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なIDです"})
		}

		todo, err := h.todoUsecase.GetTodoByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Todoが見つかりません"})
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (h *todoHandler) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なIDです"})
		}

		var req UpdateTodoRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なリクエストです"})
		}

		todo, err := h.todoUsecase.UpdateTodo(uint(id), req.Title, req.Completed)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (h *todoHandler) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "無効なIDです"})
		}

		if err := h.todoUsecase.DeleteTodo(uint(id)); err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Todoが見つかりません", "id": id})
		}

		return c.JSON(http.StatusOK, echo.Map{"message": "Todoが削除されました", "id": id})
	}
}
