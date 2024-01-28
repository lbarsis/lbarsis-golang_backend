package main

import (
	"go-backend/internal/handler"
	"go-backend/internal/model"
	"go-backend/pkg/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConnectToDB()
	config.DB.AutoMigrate(&model.User{})

	// Register routes
	e.POST("/users", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}
