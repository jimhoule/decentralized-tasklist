package payloads

type GetTaskPayload struct {
	Id                     string
	EncodedContractAddress string
	EncodedPrivateKey      string
}