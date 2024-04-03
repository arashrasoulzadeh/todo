package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
)

type TaskDatabase struct {
	db models.Database
}

func NewTaskRepository() *TaskDatabase {
	return &TaskDatabase{
		db: models.Database{},
	}
}

type TaskRepository interface {
	TaskList() models.Task
	CreateTask(task models.Task) models.Task
	UpdateTask(id int, task models.Task) models.Task
}

func (t *TaskDatabase) TaskList() []models.Task {
	return t.db.Items
}

func (t *TaskDatabase) CreateTask(title string, desc string) models.Task {
	task := models.Task{
		Id:     len(t.db.Items) + 1,
		Title:  title,
		Desc:   sql.NullString{String: desc, Valid: true},
		Status: 0,
	}
	t.db.Items = append(t.db.Items, task)

	return task
}
func (t *TaskDatabase) UpdateTask(id int, title string, desc string, status uint8) (*models.Task, error) {
	for index, item := range t.db.Items {
		if item.Id == id {
			task := models.Task{
				Id:     id,
				Title:  title,
				Desc:   sql.NullString{String: desc, Valid: true},
				Status: status,
			}
			t.db.Items[index] = task
			return &task, nil
		}
	}
	return nil, errors.New("item not found")
}
