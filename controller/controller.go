package controller

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"
)

// Usecase - Usecase interface
type Usecase interface {
	GetTokenCache(key string) (string, error)
}

// Controller - Controller struct
type Controller struct {
	u Usecase
}

// NewController - Creates a new controller
func NewController(u Usecase) *Controller {
	return &Controller{u}
}

// IndexRoute - Index route
func (c Controller) IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome")
}

// GetToken - Get token from cache
func (c *Controller) GetToken(w http.ResponseWriter, r *http.Request) {
	// Get headers or params from request
	// Data provided by GAuth
	clientID := r.Header.Get("client_id")
	clientSecretValue := r.Header.Get("client_secret")
	// Encryption
	result := b64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecretValue))
	fmt.Println(result)
	token, err := c.u.GetTokenCache(result)

	if err != nil {
		fmt.Println("error controller")
		fmt.Println(err)
	}
	fmt.Println("Result token: ", token)
}
