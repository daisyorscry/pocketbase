package service

import (
	data "belajar/Data"
	repository "belajar/Repositories"

	"github.com/pocketbase/pocketbase/models"
)

type TodoServiceImplementation struct {
	Repo repository.TodolistRepository
}

func NewTodolistService(repo repository.TodolistRepository) TodolistService {
	return &TodoServiceImplementation{
		Repo: repo,
	}
}

func (s *TodoServiceImplementation) GetAllTodolists() ([]*models.Record, error) {
	return s.Repo.GetAllTodolists()
}

func (s *TodoServiceImplementation) GetTodolistById(id string) (*models.Record, error) {
	return s.Repo.GetTodolistById(id)
}
func (s *TodoServiceImplementation) GetTodolistByName(name string) (*models.Record, error) {
	return s.Repo.GetTodolistByName(name)
}
func (s *TodoServiceImplementation) CreateTodolist(todo *data.Todo) (*models.Record, error) {
	return s.Repo.CreateTodolist(todo)
}
func (s *TodoServiceImplementation) UpdateTodolistById(todo *data.Todo) (*models.Record, error) {
	return s.Repo.UpdateTodolistById(todo)
}
func (s *TodoServiceImplementation) DeleteTodolistById(id string) error {
	return s.Repo.DeleteTodolistById(id)
}
