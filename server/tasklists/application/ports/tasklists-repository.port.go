package ports

import "tasklist/domain/models"

type TasklistsRepositoryPort interface {
	Get(encodedContractAddress string, encodedPrivateKey string) (*models.Tasklist, error)
	Create(tasklist *models.Tasklist, encodedPrivateKey string) (*models.Tasklist, error)
	GetTask(id string, encodedContractAddress string, encodedPrivateKey string) (*models.Task, error)
	AddTask(task *models.Task, encodedContractAddress string, encodedPrivateKey string) (*models.Task, error)
	UpdateTask(task *models.Task, encodedContractAddress string, encodedPrivateKey string) (*models.Task, error)
	DeleteTask(id string, encodedContractAddress string, encodedPrivateKey string) (string, error)
}