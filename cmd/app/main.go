package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"main.go/internal/database"
	"main.go/internal/handlers"
	"main.go/internal/taskService"
	"main.go/internal/userService"
	"main.go/internal/web/tasks"
	"main.go/internal/web/users"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}, &userService.User{}); err != nil {
		fmt.Println("No automigrate")
	}

	taskRepo := taskService.NewTaskRepository(database.DB)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUsersHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Request().Header.Set("Content-Type", "application/json")
			fmt.Println("Request Content-Type:", c.Request().Header.Get("Content-Type"))
			return next(c)
		}
	})

	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
