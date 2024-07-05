package factories

import "tasklist/domain/models"

type TasklistsFactory struct{}

func (tf *TasklistsFactory) Create(name string, description string, email string) *models.Tasklist {
	return &models.Tasklist{
		Name: name,
		Description: description,
		Email: email,
		Tasks: []*models.Task{},
	}
}