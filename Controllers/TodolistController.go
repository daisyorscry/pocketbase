package controller

import (
	"github.com/labstack/echo/v5"
)

type TodolistController interface {
	GetAllTodolists(ctx echo.Context) error
	GetTodolistById(ctx echo.Context) error
	GetTodolistByName(ctx echo.Context) error
	CreateTodolist(ctx echo.Context) error
	UpdateTodolistById(ctx echo.Context) error
	// UpdateTodolistByName(ctx echo.Context) error
	// UpdateTodolistByEmail(ctx echo.Context)error
	DeleteTodolistById(ctx echo.Context) error
	// DeleteTodolistByName(ctx echo.Context) error
	// DeleteTodolistByEmail(ctx echo.Context) error
}
