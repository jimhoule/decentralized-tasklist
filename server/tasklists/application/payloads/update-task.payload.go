package payloads

type UpdateTaskPayload struct {
	Id                     string
	Name                   string
	Description            string
	IsDone                 bool
	EncodedContractAddress string
	EncodedPrivateKey      string
}