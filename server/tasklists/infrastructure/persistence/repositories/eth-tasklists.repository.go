package repositories

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"tasklist/domain/models"
	tasklistContractBinder "tasklist/infrastructure/persistence/repositories/contracts/tasklist"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthTasklistsRepository struct {
	EthClient *ethclient.Client
}

func (etr *EthTasklistsRepository) Get(encodedContractAddress string, encodedPrivateKey string) (*models.Tasklist, error) {
	// Gets private key and public address
	_, publicAddress, err := extractCredentials(encodedPrivateKey)
	if err != nil {
		return nil, err
	}

	// Gets tasklist contract interactor (Allows us to interact with deployed contract)
	contractAddress := common.HexToAddress(encodedContractAddress)

	tasklistContract, err := tasklistContractBinder.NewTasklistCaller(
		contractAddress,
		etr.EthClient,
	)
	if err != nil {
		return nil, err
	}

	// Sets caller options
	callerOptions := &bind.CallOpts{
		From: publicAddress,
	}

	// Gets tasklist info
	tasklistInfo, err := tasklistContract.GetInfo(callerOptions)
	if err != nil {
		return nil, err
	}

	// Gets tasks
	contractTasks, err := tasklistContract.GetTasks(callerOptions)
	if err != nil {
		return nil, err
	}

	tasks := []*models.Task{}
	for _, contractTask := range contractTasks { 
		task := &models.Task{
			Id: contractTask.Id,
			Name: contractTask.Name,
			IsDone: contractTask.IsDone,
		}

		tasks = append(tasks, task)
	}

	tasklist := &models.Tasklist{
		Name: tasklistInfo.Name,
		Description: tasklistInfo.Description,
		Email: tasklistInfo.Email,
		Tasks: tasks,
	}


	return tasklist, nil
}

func (etr *EthTasklistsRepository) Create(tasklist *models.Tasklist, encodedPrivateKey string) (*models.Tasklist, error) {
	// Gets private key and public address
	privateKey, publicAddress, err := extractCredentials(encodedPrivateKey)
	if err != nil {
		return nil, err
	}

	// Gets nonce that should be used for next transaction
	nonce, err := etr.EthClient.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		return nil, err
	}

	// Deploys tasklist contract
	transactorOptions, err := getTransactorOptions(privateKey, etr.EthClient)
	if err != nil {
		return nil, err
	}
	transactorOptions.Nonce = big.NewInt(int64(nonce))

	_, _, _, err = tasklistContractBinder.DeployTasklist(
		transactorOptions,
		etr.EthClient,
		tasklist.Name,
		tasklist.Description,
		tasklist.Email,
	)
	if err != nil {
		return nil, err
	}

	return tasklist, nil
}

func (etr *EthTasklistsRepository) GetTask(id string, encodedContractAddress string, encodedPrivateKey string) (*models.Task, error) {
	// Gets public address
	_, publicAddress, err := extractCredentials(encodedPrivateKey)
	if err != nil {
		return nil, err
	}

	// Gets tasklist contract interactor (Allows us to interact with deployed contract)
	contractAddress := common.HexToAddress(encodedContractAddress)

	tasklistContract, err := tasklistContractBinder.NewTasklistCaller(
		contractAddress,
		etr.EthClient,
	)
	if err != nil {
		return nil, err
	}

	// Gets task
	callerOptions := &bind.CallOpts{
		From: publicAddress,
	}

	contractTask, err := tasklistContract.GetTask(callerOptions, id)
	if err != nil {
		return nil, err
	}

	task := &models.Task{
		Id: contractTask.Id,
		Name: contractTask.Name,
		Description: contractTask.Description,
		IsDone: contractTask.IsDone,
	}


	return task, nil
}

func (etr *EthTasklistsRepository) AddTask(task *models.Task, encodedContractAddress string, encodedPrivateKey string) (*models.Task, error) {
	// Gets private key and public address
	privateKey, _, err := extractCredentials(encodedPrivateKey)
	if err != nil {
		return nil, err
	}

	// Gets tasklist contract interactor (Allows us to interact with deployed contract)
	contractAddress := common.HexToAddress(encodedContractAddress)

	tasklistContract, err := tasklistContractBinder.NewTasklistTransactor(
		contractAddress,
		etr.EthClient,
	)
	if err != nil {
		return nil, err
	}

	// Adds task
	transactorOptions, err := getTransactorOptions(privateKey, etr.EthClient)
	if err != nil {
		return nil, err
	}

	_, err = tasklistContract.AddTask(transactorOptions, task.Id, task.Name, task.Description)
	if err != nil {
		return nil, err
	}


	return task, nil
}

func (etr *EthTasklistsRepository) UpdateTask(task *models.Task, encodedContractAddress string, encodedPrivateKey string) (*models.Task, error) {
	// Gets private key and public address
	privateKey, _, err := extractCredentials(encodedPrivateKey)
	if err != nil {
		return nil, err
	}

	// Gets tasklist contract interactor
	contractAddress := common.HexToAddress(encodedContractAddress)

	tasklistContract, err := tasklistContractBinder.NewTasklistTransactor(
		contractAddress,
		etr.EthClient,
	)
	if err != nil {
		return nil, err
	}

	// Updates task
	transactorOptions, err := getTransactorOptions(privateKey, etr.EthClient)
	if err != nil {
		return nil, err
	}

	_, err = tasklistContract.UpdateTask(
		transactorOptions,
		task.Id,
		task.Name,
		task.Description,
		task.IsDone,
	)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (etr *EthTasklistsRepository) DeleteTask(id string, encodedContractAddress string, encodedPrivateKey string) (string, error) {
	// Gets private key and public address
	privateKey, _, err := extractCredentials(encodedPrivateKey)
	if err != nil {
		return "", err
	}

	// Gets tasklist contract interactor (Allows us to interact with deployed contract)
	contractAddress := common.HexToAddress(encodedContractAddress)

	tasklistContract, err := tasklistContractBinder.NewTasklistTransactor(
		contractAddress,
		etr.EthClient,
	)
	if err != nil {
		return "", err
	}

	// Deletes task
	transactorOptions, err := getTransactorOptions(privateKey, etr.EthClient)
	if err != nil {
		return "", err
	}

	_, err = tasklistContract.DeleteTask(transactorOptions, id)
	if err != nil {
		return "", err
	}


	return id, nil
}

// NOTE: Gets private key and public address from encoded private key
func extractCredentials(encodedPrivateKey string) (*ecdsa.PrivateKey, common.Address, error) {
	// Gets private key
	privateKeyData, err := hexutil.Decode(encodedPrivateKey)
	if err != nil {
		return nil, common.Address{}, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyData)
	if err != nil {
		return nil, common.Address{}, err
	}

	// Gets public address
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	return privateKey, publicAddress, nil
}

// NOTE: Gets transactor options with basic settings
func getTransactorOptions(privateKey *ecdsa.PrivateKey, ethClient *ethclient.Client) (*bind.TransactOpts, error) {
	// Gets gas price (transaction fees)
	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	// Gets blockchain ID
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	// Deletes task
	transactorOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	transactorOptions.GasPrice = gasPrice
	transactorOptions.GasLimit = uint64(3000000)

	return transactorOptions, nil
}