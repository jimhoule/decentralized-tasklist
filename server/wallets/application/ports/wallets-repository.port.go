package ports

import "wallet/domain/models"

type WalletsRepositoryPort interface {
	Create() (*models.Wallet, error)
	CreateKeystore(password string) (*models.Wallet, error)
}