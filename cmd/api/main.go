package main

import (
	"fmt"
	config "go-tutorial/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// データベース接続
	_, err := config.InitDB()
	if err != nil {
		panic("データベース接続に失敗しました: " + err.Error())
	}
	fmt.Println("データベース接続に成功しました！")

	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// サーバーの起動
	e.Logger.Fatal(e.Start(":8080"))
}
