package payloads

type DeleteTaskPayload struct {
	Id                     string
	EncodedContractAddress string
	EncodedPrivateKey      string
}