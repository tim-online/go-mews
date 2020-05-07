package json

type BaseRequest struct {
	AccessToken  string `json:"AccessToken"`
	ClientToken  string `json:"ClientToken,omitempty"`
	LanguageCode string `json:"LanguageCode,omitempty"`
	CultureCode  string `json:"CultureCode,omitempty"`
	Client       string `json:"Client,omitempty"`
}

func (req *BaseRequest) SetAccessToken(token string) {
	req.AccessToken = token
}

func (req *BaseRequest) SetClientToken(token string) {
	req.ClientToken = token
}

func (req *BaseRequest) SetLanguageCode(code string) {
	req.LanguageCode = code
}

func (req *BaseRequest) SetCultureCode(code string) {
	req.CultureCode = code
}
