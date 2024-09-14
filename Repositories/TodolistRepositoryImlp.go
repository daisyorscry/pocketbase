package repository

import (
	data "belajar/Data"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

// struct for implementation repository
type TodoRepositoryImplementation struct {
	App *pocketbase.PocketBase
}

func NewTodolistRepository(app *pocketbase.PocketBase) TodolistRepository {
	return &TodoRepositoryImplementation{
		App: app,
	}
}

// get all todolist implementation
func (r *TodoRepositoryImplementation) GetAllTodolists() ([]*models.Record, error) {
	records, err := r.App.Dao().FindRecordsByExpr("todolist", nil)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r *TodoRepositoryImplementation) GetTodolistById(id string) (*models.Record, error) {
	record, err := r.App.Dao().FindRecordById("todolist", id)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (r *TodoRepositoryImplementation) GetTodolistByName(name string) (*models.Record, error) {
	record, err := r.App.Dao().FindRecordsByExpr("todolist", dbx.NewExp("name = {:name}", dbx.Params{"name": name}))
	if err != nil || len(record) == 0 {
		return nil, err
	}

	return record[0], nil
}

func (r *TodoRepositoryImplementation) CreateTodolist(todo *data.Todo) (*models.Record, error) {
	collection, err := r.App.Dao().FindCollectionByNameOrId("todolist")
	if err != nil {
		return nil, err
	}
	record := models.NewRecord(collection)
	record.Set("title", todo.Title)
	record.Set("completed", todo.Completed)

	if err := r.App.Dao().SaveRecord(record); err != nil {
		return nil, err
	}

	return record, err

}

func (r *TodoRepositoryImplementation) UpdateTodolistById(todo *data.Todo) (*models.Record, error) {
	record, err := r.App.Dao().FindRecordById("todolist", todo.Id)
	if err != nil {
		return nil, err
	}
	record.Set("title", todo.Title)
	record.Set("completed", todo.Completed)

	if err := r.App.Dao().SaveRecord(record); err != nil {
		return nil, err
	}

	return record, err
}

func (r *TodoRepositoryImplementation) DeleteTodolistById(id string) error {
	record, err := r.App.Dao().FindRecordById("todolist", id)
	if err != nil {
		return err
	}

	return r.App.Dao().DeleteRecord(record)
}
