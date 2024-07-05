package services

import (
	"tasklist/application/payloads"
	"tasklist/application/ports"
	"tasklist/domain/factories"
	"tasklist/domain/models"
)

type TasklistsService struct {
	TasksFactory *factories.TasksFactory
	TasklistsFactory *factories.TasklistsFactory
	TasklistsRepository ports.TasklistsRepositoryPort
}

func (ts *TasklistsService) Get(getTasklistPayload * payloads.GetTasklistPayload) (*models.Tasklist, error) {
	return ts.TasklistsRepository.Get(getTasklistPayload.EncodedContractAddress, getTasklistPayload.EncodedPrivateKey)
}

func (ts *TasklistsService) Create(createTasklistPayload *payloads.CreateTasklistPayload) (*models.Tasklist, error) {
	tasklist := ts.TasklistsFactory.Create(
		createTasklistPayload.Name,
		createTasklistPayload.Description,
		createTasklistPayload.Email,
	)

	return ts.TasklistsRepository.Create(tasklist, createTasklistPayload.EncodedPrivateKey)
}

func (ts *TasklistsService) GetTask(getTaskPayload * payloads.GetTaskPayload) (*models.Task, error) {
	return ts.TasklistsRepository.GetTask(
		getTaskPayload.Id,
		getTaskPayload.EncodedContractAddress,
		getTaskPayload.EncodedPrivateKey,
	)
}

func (ts *TasklistsService) AddTask(addTaskPayload *payloads.AddTaskPayload) (*models.Task, error) {
	task := ts.TasksFactory.Create(addTaskPayload.Name, addTaskPayload.Description)

	return ts.TasklistsRepository.AddTask(
		task,
		addTaskPayload.EncodedContractAddress,
		addTaskPayload.EncodedPrivateKey,
	)
}

func (ts *TasklistsService) UpdateTask(updateTaskPayload *payloads.UpdateTaskPayload) (*models.Task, error) {
	task := &models.Task{
		Id: updateTaskPayload.Id,
		Name: updateTaskPayload.Name,
		Description: updateTaskPayload.Description,
		IsDone: updateTaskPayload.IsDone,
	}

	return ts.TasklistsRepository.UpdateTask(
		task,
		updateTaskPayload.EncodedContractAddress,
		updateTaskPayload.EncodedPrivateKey,
	)
}

func (ts *TasklistsService) DeleteTask(deleteTaskPayload *payloads.DeleteTaskPayload) (string, error) {
	return ts.TasklistsRepository.DeleteTask(
		deleteTaskPayload.Id,
		deleteTaskPayload.EncodedContractAddress,
		deleteTaskPayload.EncodedPrivateKey,
	)
}