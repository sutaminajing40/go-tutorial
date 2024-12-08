package main

import (
	"fmt"
	"go-tutorial/backend/internal/handler"
	"go-tutorial/backend/internal/repository"
	"go-tutorial/backend/internal/usecase"
	config "go-tutorial/backend/pkg/database"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 環境変数の読み込み
	env := os.Getenv("GO_ENV")
	if env == "" || env == "local" {
		// 本番環境用の.envを先に試す
		err := godotenv.Load(".env")
		if err != nil {
			// 開発環境用の.env.localを試す
			if err := godotenv.Load(".env.local"); err != nil {
				fmt.Println("Warning: No .env or .env.local file found")
			}
		}
	}

	// データベース接続
	db, err := config.InitDB()
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
	port := os.Getenv("SERVER_PORT")
	fmt.Printf("サーバーを起動します... ポート: %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}
