package payloads

type CreateTasklistPayload struct {
	Name              string
	Description       string
	Email             string
	EncodedPrivateKey string
}