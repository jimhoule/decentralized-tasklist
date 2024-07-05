package main

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"wallet/application/services"
	"wallet/infrastructure/persistence/repositories"
	"wallet/presenters/rpc/handlers"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	ethClient, err := ethclient.DialContext(context.Background(), os.Getenv("ETH_RPC_URL"))
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	walletsRpcHandler := &handlers.WalletsRpcHandler{
		WalletsService: &services.WalletsService{
			WalletsRepository: &repositories.EthWalletsRepository{
				EthClient: ethClient,
			},
		},
	}

	// Registers Rpc Server
	rpc.Register(walletsRpcHandler)

	// Starts Rpc Server
	listener, err := net.Listen(os.Getenv("RPC_NETWORK"), fmt.Sprintf("%s:%s", os.Getenv("RPC_ADDRESS"), os.Getenv("RPC_PORT")))
	if err != nil {
		fmt.Println("Error registering Rpc Server in Wallets microservice")
	}
	defer listener.Close();

	for {
		rpcConnection, err := listener.Accept()
		if err != nil {
			continue;
		}

		go rpc.ServeConn(rpcConnection)
	}
}