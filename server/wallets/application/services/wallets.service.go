package services

import (
	"wallet/application/ports"
	"wallet/domain/models"
)

type WalletsService struct {
	WalletsRepository ports.WalletsRepositoryPort
}

func (ws *WalletsService) Create() (*models.Wallet, error) {
	return ws.WalletsRepository.Create()
}