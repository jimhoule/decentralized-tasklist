package factories

import (
	"tasklist/domain/models"

	"github.com/google/uuid"
)

type TasksFactory struct{}

func (tf *TasksFactory) Create(name string, description string) *models.Task {
	return &models.Task{
		Id: uuid.NewString(),
		Name: name,
		Description: description,
		IsDone: false,
	}
}