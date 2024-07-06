package models

type Wallet struct {
	PrivateKey    string `json:"privateKey"`
	PublicKey     string `json:"publicKey"`
	PublicAddress string `json:"publicAddress"`
}