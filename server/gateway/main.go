package main

import (
	"fmt"
	"gateway/presenters/http/controllers"
	"gateway/presenters/http/router"
	"log"
	"net/http"
	"net/rpc"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// Starts connection to tasklists microservice
	tasklistsRpcClient := createRpcClient(
		os.Getenv("TASKLISTS_RPC_NETWORK"),
		fmt.Sprintf("%s:%s", os.Getenv("TASKLISTS_RPC_ADDRESS"), os.Getenv("TASKLISTS_RPC_PORT")),
	)
	defer tasklistsRpcClient.Close()

	// Starts connection to wallets microservice
	walletsRpcClient := createRpcClient(
		os.Getenv("WALLETS_RPC_NETWORK"),
		fmt.Sprintf("%s:%s", os.Getenv("WALLETS_RPC_ADDRESS"), os.Getenv("WALLETS_RPC_PORT")),
	)
	defer walletsRpcClient.Close()

	// Gets main router
	mainRouter := router.Get()

	// Creates tasklists routes
	taskListsController := controllers.TaskListsController{
		TasklistsRpcClient: tasklistsRpcClient,
	}
	mainRouter.Get("/tasklists/{address}/{key}", taskListsController.Get)
	mainRouter.Post("/tasklists/{key}", taskListsController.Create)
	mainRouter.Get("/tasklists/tasks/{id}/{address}/{key}", taskListsController.GetTask)
	mainRouter.Post("/tasklists/tasks/{address}/{key}", taskListsController.AddTask)
	mainRouter.Put("/tasklists/tasks/{id}/{address}/{key}", taskListsController.UpdateTask)
	mainRouter.Delete("/tasklists/tasks/{id}/{address}/{key}", taskListsController.DeleteTask)

	// Creates wallets routes
	walletsController := controllers.WalletsController{
		WalletsRpcClient: walletsRpcClient,
	}
	mainRouter.Post("/wallets", walletsController.Create)

	// Creates server
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("HTTP_URL"), os.Getenv("HTTP_PORT")),
		Handler: mainRouter,
	}

	// Starts server
	err = server.ListenAndServe();
	if err != nil {
		log.Panic(err);
	}
}

func createRpcClient(network string, address string) *rpc.Client {
	rpcClient, err := rpc.Dial(network, address)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	return rpcClient
}