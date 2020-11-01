package model

import (
	"encoding/json"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewError(err error) []byte {
	resp := ErrorResponse{Error: err.Error()}
	bytes, _ := json.Marshal(&resp)
	return bytes
}
