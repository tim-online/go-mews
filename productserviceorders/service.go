package productserviceorders

import "github.com/tim-online/go-mews/json"

type APIService struct {
	Client *json.Client
}

func NewAPIService() *APIService {
	return &APIService{}
}

