package common

type PublishInput struct {
	Action   string            `json:"action"`
	Metadata map[string]string `json:"metadata"`
	Request  interface{}       `json:"request"`
	Response interface{}       `json:"response"`
	_        struct{}
}

type PublishOutput struct {
	_ struct{}
}
