package args

type UpdateTaskArgs struct {
	Id                     string
	Name                   string
	Description            string
	IsDone                 bool
	EncodedContractAddress string
	EncodedPrivateKey      string
}