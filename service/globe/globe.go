package service

import (
	"encoding/json"
	"fmt"
	"tokens/model"

	cError "github.com/coreos/etcd/error"
	"gopkg.in/resty.v1"
)

const globeEndpoint = "/v1/oauth/accessToken"

// Service - Service struct
type Service struct {
	globeClient *resty.Client
}

// NewGlobeClient - Creates a new Globe ckient
func NewGlobeClient(host, timeout string) *Service {
	client := resty.New().
		SetHostURL(host).
		//SetTimeout("5s").
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			if r.IsSuccess() {
				return nil
			}

			return cError.NewError(r.StatusCode(), "error", 0)
		})
	return &Service{globeClient: client}
}

// GetToken - Get token from Globe
func (s *Service) GetToken(key string) (string, error) {
	fmt.Println("Globe GetToken")
	token := &model.Token{
		DeveloperEmail: "test@gmail.com",
		TokenType:      "BearerToken",
		IssuedAt:       "1623974364",
		AccessToken:    "KjfP8EAGfo4DR6HdeWKywra7poeU",
		ExpiresIn:      "60",
		Status:         "approved",
	}
	// `{"DeveloperEmail":"test@gmail.com",` +
	// 	`"TokenType":"BearerToken",` +
	// 	`"IssuedAt":"1623974364",` +
	// 	`"AccessToken":"KjfP8EAGfo4DR6HdeWKywra7poeU",` +
	// 	`"ExpiresIn":"60",` +
	// 	`"Status":"approved"` +
	// 	`}`

	t, _ := json.Marshal(&token)

	return string(t), nil
}
