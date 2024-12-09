package main

import (
	env "backend/config"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/usecase"
	dbconfig "backend/pkg/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 設定の読み込み
	cfg := env.Load()

	// データベース接続
	db, err := dbconfig.InitDB()
	if err != nil {
		panic("データベース接続に失敗しました: " + err.Error())
	}
	fmt.Println("データベース接続に成功しました！")

	// 依存関係の注入
	todoRepo := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // CORSを有効化

	// ルーティングの設定
	api := e.Group("/api")
	{
		todos := api.Group("/todos")
		todos.POST("", todoHandler.CreateTodo())
		todos.GET("", todoHandler.GetAllTodos())
		todos.GET("/:id", todoHandler.GetTodoByID())
		todos.PUT("/:id", todoHandler.UpdateTodo())
		todos.DELETE("/:id", todoHandler.DeleteTodo())
	}

	// サーバーの起動
	port := cfg.Port
	fmt.Printf("サーバーを起動します... ポート: %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}
