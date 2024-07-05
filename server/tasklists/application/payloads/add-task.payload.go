package payloads

type AddTaskPayload struct {
	Name                   string
	Description            string
	EncodedContractAddress string
	EncodedPrivateKey      string
}