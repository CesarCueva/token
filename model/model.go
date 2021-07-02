package model

// Token - Token structure
type Token struct {
	DeveloperEmail string `json:"developerEmail"`
	TokenType      string `json:"tokenType"`
	IssuedAt       string `json:"issuedAt"`
	AccessToken    string `json:"accessToken"`
	ExpiresIn      string `json:"expiresIn"`
	Status         string `json:"status"`
}
