package server

import "encoding/json"

type APIErrors struct {
	ErrorType string `json:"error_type"`
	ErrorMsg  string `json:"error_msg"`
}

func (a APIErrors) Error() string {
	bs, _ := json.Marshal(a)
	return string(bs)
}

const (
	ErrInvalidAccessToken = "invalid access token"
	ErrDataNotFound = "data not found"
	ErrLogin = "login with token error"
	ErrTokenValidation = "provided token is not valid"
)
