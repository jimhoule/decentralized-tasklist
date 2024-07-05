package controllers

import (
	"gateway/domain/models"
	"gateway/presenters/http/dtos"
	"gateway/presenters/http/router"
	"gateway/presenters/http/utils"
	"net/http"
	"net/rpc"
)

type TaskListsController struct{
	TasklistsRpcClient *rpc.Client
}

func (tc *TaskListsController) Get(writer http.ResponseWriter, request *http.Request) {
	encodedContractAddress := router.GetUrlParam(request, "address")
	encodedPrivateKey := router.GetUrlParam(request, "key")

	getTasklistArgs := struct{  
		EncodedContractAddress string
		EncodedPrivateKey string
	}{
		EncodedContractAddress: encodedContractAddress,
		EncodedPrivateKey: encodedPrivateKey,
	}

	var tasklist models.Tasklist
	// NOTE: TasklistsRpcHandler is the registered Rpc Server in tasklists microservice
	err := tc.TasklistsRpcClient.Call("TasklistsRpcHandler.Get", getTasklistArgs, &tasklist);
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, tasklist)
}

func (tc *TaskListsController) Create(writer http.ResponseWriter, request *http.Request) {
	encodedPrivateKey := router.GetUrlParam(request, "key")

	var createTasklistDto dtos.CreateTasklistDto
	err := utils.ReadHttpRequestBody(writer, request, &createTasklistDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	createTasklistArgs := struct{  
		Name string
		Description string
		Email string
		EncodedPrivateKey string
	}{
		Name: createTasklistDto.Name,
		Description: createTasklistDto.Description,
		Email: createTasklistDto.Email,
		EncodedPrivateKey: encodedPrivateKey,
	}

	var tasklist models.Tasklist;
	err = tc.TasklistsRpcClient.Call("TasklistsRpcHandler.Create", createTasklistArgs, &tasklist);
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusCreated, tasklist)
}

func (tc *TaskListsController) GetTask(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")
	encodedContractAddress := router.GetUrlParam(request, "address")
	encodedPrivateKey := router.GetUrlParam(request, "key")

	getTaskArgs := struct{
		Id string
		EncodedContractAddress string
		EncodedPrivateKey string
	}{
		Id: id,
		EncodedContractAddress: encodedContractAddress,
		EncodedPrivateKey: encodedPrivateKey,
	}

	var task models.Task
	// NOTE: TasklistsRpcHandler is the registered Rpc Server in tasklists microservice
	err := tc.TasklistsRpcClient.Call("TasklistsRpcHandler.GetTask", getTaskArgs, &task);
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, task)
}

func (tc *TaskListsController) AddTask(writer http.ResponseWriter, request *http.Request) {
	encodedContractAddress := router.GetUrlParam(request, "address")
	encodedPrivateKey := router.GetUrlParam(request, "key")

	var addTaskDto dtos.AddTaskDto
	err := utils.ReadHttpRequestBody(writer, request, &addTaskDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	addTaskArgs := struct{  
		Name string
		Description string
		EncodedContractAddress string
		EncodedPrivateKey string
	}{
		Name: addTaskDto.Name,
		Description: addTaskDto.Description,
		EncodedContractAddress: encodedContractAddress,
		EncodedPrivateKey: encodedPrivateKey,
	}

	var task models.Task;
	err = tc.TasklistsRpcClient.Call("TasklistsRpcHandler.AddTask", addTaskArgs, &task);
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, task)
}

func (tc *TaskListsController) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")
	encodedContractAddress := router.GetUrlParam(request, "address")
	encodedPrivateKey := router.GetUrlParam(request, "key")

	var updateTaskDto dtos.UpdateTaskDto
	err := utils.ReadHttpRequestBody(writer, request, &updateTaskDto)
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	updateTaskArgs := struct{  
		Id string
		Name string
		Description string
		IsDone bool
		EncodedContractAddress string
		EncodedPrivateKey string
	}{
		Id: id,
		Name: updateTaskDto.Name,
		Description: updateTaskDto.Description,
		IsDone: updateTaskDto.IsDone,
		EncodedContractAddress: encodedContractAddress,
		EncodedPrivateKey: encodedPrivateKey,
	}

	var task models.Task;
	err = tc.TasklistsRpcClient.Call("TasklistsRpcHandler.UpdateTask", updateTaskArgs, &task);
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusOK, task)
}

func (tc *TaskListsController) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	id := router.GetUrlParam(request, "id")
	encodedContractAddress := router.GetUrlParam(request, "address")
	encodedPrivateKey := router.GetUrlParam(request, "key")

	deleteTaskArgs := struct{  
		Id string
		EncodedContractAddress string
		EncodedPrivateKey string
	}{
		Id: id,
		EncodedContractAddress: encodedContractAddress,
		EncodedPrivateKey: encodedPrivateKey,
	}

	err := tc.TasklistsRpcClient.Call("TasklistsRpcHandler.DeleteTask", deleteTaskArgs, nil);
	if err != nil {
		utils.WriteHttpError(writer, http.StatusInternalServerError, err)
		return
	}

	utils.WriteHttpResponse(writer, http.StatusNoContent, nil)
}