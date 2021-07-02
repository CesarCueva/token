package usecase

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"tokens/model"

	cError "github.com/coreos/etcd/error"
)

// Redis - Redis interface
// type Redis interface {
// 	GetToken(key string) (string, error)
// 	SetToken(key, token string, expiration time.Time) error
// }

// Cache - Cache interface
type Cache interface {
	GetToken(key string) (interface{}, error)
	SetToken(key, token string) error
}

// Globe - Globe interface
type Globe interface {
	GetToken(key string) (string, error)
}

// Usecase - Usecase struct
type Usecase struct {
	cacheClient Cache
	globeClient Globe
}

// NewUsecase -
func NewUsecase(c Cache, g Globe) *Usecase {
	return &Usecase{cacheClient: c, globeClient: g}
}

// GetTokenCache - Get token from Cache
func (u *Usecase) GetTokenCache(key string) (string, error) {
	tempToken, err := u.cacheClient.GetToken(key)
	if err == nil {
		// fmt.Println("1: ", err)
		// fmt.Println("temp: ", tempToken)
		return fmt.Sprintf("%v", tempToken), err
	}

	isNewToken := false

	if tempToken == nil {
		fmt.Println("Empty token")
		tempToken, err = u.globeClient.GetToken(key)
		if err != nil {
			//fmt.Println("2")
			return "", err
		}
		isNewToken = true
	}

	token, tokenExpiration, err := validateToken(fmt.Sprintf("%v", tempToken))
	if err != nil {
		//fmt.Println("3")
		return "", err
	}

	if !isNewToken && time.Now().Unix() > tokenExpiration.Unix() {
		fmt.Println("Token expired")
		// call globe to get a token
		if isNewToken {
			return "", cError.NewError(500, "Token created is already expired...", 1)
		}
		newToken, err := u.globeClient.GetToken(key)
		if err != nil {
			//fmt.Println("4")
			return "", err
		}

		token, tokenExpiration, err = validateToken(newToken)
		if err != nil {
			//fmt.Println("5")
			return "", err
		}

		err = u.cacheClient.SetToken(key, newToken)
		if err != nil {
			//fmt.Println("6")
			return "", err
		}

		return fmt.Sprintf("%v", token), nil
	}
	t, err := json.Marshal(&token)
	if err != nil {
		//fmt.Println("7")
		return "", err
	}

	err = u.cacheClient.SetToken(key, string(t))
	if err != nil {
		//fmt.Println("8")
		return "", err
	}

	return fmt.Sprintf("%v", token), nil
}

func validateToken(tempToken string) (*model.Token, *time.Time, error) {
	// String to []bytes
	tokenBytes := []byte(tempToken)

	// Unmarshal to token struct
	token := &model.Token{}
	err := json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, nil, err
	}

	// Validate token expiration
	expiration, _ := strconv.ParseInt(token.ExpiresIn, 10, 64)
	issuedAt, _ := strconv.ParseInt(token.IssuedAt, 10, 64)
	issuedDate := time.Unix(issuedAt, 0)
	expirationDate := issuedDate.Add(time.Second * time.Duration(expiration))
	fmt.Println("issued: ", issuedDate)
	fmt.Println("timein: ", expirationDate)
	return token, &expirationDate, nil
}
