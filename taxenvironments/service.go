package taxenvironments

import "github.com/tim-online/go-mews/json"

type Service struct {
	Client *json.Client
}

func NewService() *Service {
	return &Service{}
}
