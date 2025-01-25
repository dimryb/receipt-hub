package main

import (
	"fmt"
	swagger "github.com/swaggo/http-swagger" // подключаем Swagger
	_ "receipt-loader/docs"
	"receipt-loader/internal/app"
)

// @title Receipt Hub API
// @version 1.0
// @description This is a sample server for Receipt Hub.
// @host localhost:8080
// @BasePath /
func main() {
	app := app.NewApp()
	err := app.Config.Load(".env")

	if err != nil {
		panic(err)
	}

	err = app.Setup()
	if err != nil {
		panic(err)
	}

	// Настройка маршрутов Swagger UI
	app.Router.PathPrefix("/swagger/").Handler(swagger.WrapHandler)

	fmt.Println("Starting is running.")

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
