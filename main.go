package main

import (
	controller "belajar/Controllers"
	repository "belajar/Repositories"
	service "belajar/Services"
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	TodolistRepository := repository.NewTodolistRepository(app)
	TodolistService := service.NewTodolistService(TodolistRepository)
	TodolistController := controller.NewTodolistController(TodolistService)

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/todos", TodolistController.GetAllTodolists)
		e.Router.GET("/todos/id/:id", TodolistController.GetTodolistById)
		e.Router.GET("/todos/name/:name", TodolistController.GetTodolistByName)
		e.Router.POST("/todos", TodolistController.CreateTodolist)
		e.Router.PUT("/todos", TodolistController.UpdateTodolistById)
		e.Router.DELETE("/todos", TodolistController.DeleteTodolistById)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
