package handlers

import (
	"wallet/application/services"
	"wallet/domain/models"
	"wallet/presenters/rpc/args"
)

type WalletsRpcHandler struct{
	WalletsService *services.WalletsService
}

func (wrh *WalletsRpcHandler) Create(createWalletArgs *args.CreateWalletArgs, response *models.Wallet) error {
	wallet, err := wrh.WalletsService.Create()
	if err != nil {
		return err
	}

	*response = *wallet

	return nil
}