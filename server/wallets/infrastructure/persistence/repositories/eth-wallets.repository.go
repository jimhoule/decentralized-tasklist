package repositories

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"wallet/domain/models"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthWalletsRepository struct {
	EthClient *ethclient.Client
}

func (ewr *EthWalletsRepository) Create() (*models.Wallet, error) {
	// Gets private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyData := crypto.FromECDSA(privateKey)
	encodedPrivateKey := hexutil.Encode(privateKeyData)

	// Gets public key
	publicKey := privateKey.PublicKey
	publicKeyData := crypto.FromECDSAPub(&publicKey)
	encodedPublicKey := hexutil.Encode(publicKeyData)

	// Gets public address
	publicAddress := crypto.PubkeyToAddress(publicKey)
	encodedPublicAddress := publicAddress.Hex()

	wallet := &models.Wallet{
		PrivateKey: encodedPrivateKey,
		PublicKey: encodedPublicKey,
		PublicAddress: encodedPublicAddress,
	}

	return wallet, nil
}

// NOTE: Keystore is an encrypted version of the private that allows us to add an extra layer of security
func (ewr *EthWalletsRepository) CreateKeystore(password string) (*models.Wallet, error) {
	// Creates account
	key := keystore.NewKeyStore("./keystore-wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := key.NewAccount(password)
	if err != nil {
		return nil, err
	}

	// Gets private key
	bytesData, err := os.ReadFile("./keystore-wallet/keystore filename goes here")
	if err != nil {
		return nil, err
	}

	decryptedKey, err := keystore.DecryptKey(bytesData, password)
	privateData := crypto.FromECDSA(decryptedKey.PrivateKey)
	encodedPrivateKey := hexutil.Encode(privateData)

	// Gets public key
	publicData := crypto.FromECDSAPub(&decryptedKey.PrivateKey.PublicKey)
	encodedPublicKey := hexutil.Encode(publicData)

	// Gets public address
	publicAddress := crypto.PubkeyToAddress(decryptedKey.PrivateKey.PublicKey)
	encodedPublicAddress := publicAddress.Hex()

	wallet := &models.Wallet{
		PrivateKey: encodedPrivateKey,
		PublicKey: encodedPublicKey,
		PublicAddress: encodedPublicAddress,
	}

	return wallet, nil
}

func (ewr *EthWalletsRepository) GetBalance(encodedPrivateKey string) (*big.Int, error) {
	// Gets private key
	privateKeyData, err := hexutil.Decode(encodedPrivateKey)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.ToECDSA(privateKeyData)
	if err != nil {
		return nil, err
	}

	// Gets public address
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Gets balance
	balance, err := ewr.EthClient.BalanceAt(context.Background(), publicAddress, nil);
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return balance, nil
}

func (ewr *EthWalletsRepository) GetSuggestedGasPrice() *big.Int {
	// Gets gas price
	gasPrice, err := ewr.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return gasPrice
}