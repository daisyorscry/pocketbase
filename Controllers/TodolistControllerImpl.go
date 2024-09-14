package controller

import (
	data "belajar/Data"
	service "belajar/Services"
	"net/http"

	"github.com/labstack/echo/v5"
)

type TodoControllerImplementation struct {
	service service.TodolistService
}

func NewTodolistController(service service.TodolistService) TodolistController {
	return &TodoControllerImplementation{
		service: service,
	}
}

func (c *TodoControllerImplementation) GetAllTodolists(ctx echo.Context) error {
	todos, err := c.service.GetAllTodolists()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (c *TodoControllerImplementation) GetTodolistById(ctx echo.Context) error {
	id := ctx.PathParam("id")
	todo, err := c.service.GetTodolistById(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, todo)
}

func (c *TodoControllerImplementation) GetTodolistByName(ctx echo.Context) error {
	name := ctx.PathParam("name")
	todo, err := c.service.GetTodolistByName(name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, todo)
}

func (c *TodoControllerImplementation) CreateTodolist(ctx echo.Context) error {
	todo := new(data.Todo)
	if err := ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	todolist, err := c.service.CreateTodolist(todo)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, todolist)
}

func (c *TodoControllerImplementation) UpdateTodolistById(ctx echo.Context) error {
	id := ctx.PathParam("id")
	todo := new(data.Todo)
	if err := ctx.Bind(todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	todo.Id = id
	todolist, err := c.service.UpdateTodolistById(todo)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, todolist)
}

func (c *TodoControllerImplementation) DeleteTodolistById(ctx echo.Context) error {
	id := ctx.PathParam("id")
	if err := c.service.DeleteTodolistById(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
