package handlers

import (
	"tasklist/application/payloads"
	"tasklist/application/services"
	"tasklist/domain/models"
	"tasklist/presenters/rpc/args"
)

type TasklistsRpcHandler struct{
	TasklistsService *services.TasklistsService
}

func (trh *TasklistsRpcHandler) Get(getTasklistArgs *args.GetTasklistArgs, response *models.Tasklist) error {
	tasklist, err := trh.TasklistsService.Get(&payloads.GetTasklistPayload{
		EncodedContractAddress: getTasklistArgs.EncodedContractAddress,
		EncodedPrivateKey: getTasklistArgs.EncodedPrivateKey,
	})
	if err != nil {
		return err
	}

	*response = *tasklist

	return nil;
}

func (trh *TasklistsRpcHandler) Create(createTasklistArgs *args.CreateTasklistArgs, response *models.Tasklist) error {
	tasklist, err := trh.TasklistsService.Create(&payloads.CreateTasklistPayload{
		Name: createTasklistArgs.Name,
		Description: createTasklistArgs.Description,
		Email: createTasklistArgs.Email,
		EncodedPrivateKey: createTasklistArgs.EncodedPrivateKey,
	})
	if err != nil {
		return err
	}

	*response = *tasklist

	return nil;
}

func (trh *TasklistsRpcHandler) GetTask(getTaskArgs *args.GetTaskArgs, response *models.Task) error {
	task, err := trh.TasklistsService.GetTask(&payloads.GetTaskPayload{
		Id: getTaskArgs.Id,
		EncodedContractAddress: getTaskArgs.EncodedContractAddress,
		EncodedPrivateKey: getTaskArgs.EncodedPrivateKey,
	})
	if err != nil {
		return err
	}

	*response = *task

	return nil;
}

func (trh *TasklistsRpcHandler) AddTask(addTaskArgs *args.AddTaskArgs, response *models.Task) error {
	task, err := trh.TasklistsService.AddTask(&payloads.AddTaskPayload{
		Name: addTaskArgs.Name,
		Description: addTaskArgs.Description,
		EncodedContractAddress: addTaskArgs.EncodedContractAddress,
		EncodedPrivateKey: addTaskArgs.EncodedPrivateKey,
	})
	if err != nil {
		return err
	}

	*response = *task

	return nil;
}

func (trh *TasklistsRpcHandler) UpdateTask(updateTaskArgs *args.UpdateTaskArgs, response *models.Task) error {
	task, err := trh.TasklistsService.UpdateTask(&payloads.UpdateTaskPayload{
		Id: updateTaskArgs.Id,
		Name: updateTaskArgs.Name,
		Description: updateTaskArgs.Description,
		IsDone: updateTaskArgs.IsDone,
		EncodedContractAddress: updateTaskArgs.EncodedContractAddress,
		EncodedPrivateKey: updateTaskArgs.EncodedPrivateKey,
	})
	if err != nil {
		return err
	}

	*response = *task

	return nil;
}

func (trh *TasklistsRpcHandler) DeleteTask(deleteTaskArgs *args.DeleteTaskArgs, response *string) error {
	id, err := trh.TasklistsService.DeleteTask(&payloads.DeleteTaskPayload{
		Id: deleteTaskArgs.Id,
		EncodedContractAddress: deleteTaskArgs.EncodedContractAddress,
		EncodedPrivateKey: deleteTaskArgs.EncodedPrivateKey,
	})
	if err != nil {
		return err
	}

	*response = id

	return nil;
}