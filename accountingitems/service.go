package accountingitems

import "github.com/tim-online/go-mews/json"

type APIService struct {
	Client *json.Client
}

func NewService() *APIService {
	return &APIService{}
}
