package responses

import (
	"awesomeProject/internal/models"
)

type Task struct {
	Title  string `json:"title"`
	Desc   string `json:"desc,omitempty"`
	Status string `json:"status"`
}

func (t *Task) Fill(task models.Task) {
	t.Desc = task.Desc.String
	t.Title = task.Title
	switch task.Status {
	case 0:
		t.Status = "pending"
	case 1:
		t.Status = "progress"
	case 2:
		t.Status = "done"
	}
}
