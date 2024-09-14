package service

import (
	data "belajar/Data"

	"github.com/pocketbase/pocketbase/models"
)

type TodolistService interface {
	GetAllTodolists() ([]*models.Record, error)
	GetTodolistById(id string) (*models.Record, error)
	GetTodolistByName(name string) (*models.Record, error)
	CreateTodolist(todo *data.Todo) (*models.Record, error)
	UpdateTodolistById(todo *data.Todo) (*models.Record, error)
	// UpdateTodolistByName(name string) (*models.Record, error)
	// UpdateTodolistByEmail(email string) (*models.Record, error)
	DeleteTodolistById(id string) error
	// DeleteTodolistByName(name string) error
	// DeleteTodolistByEmail(email string) error
}
