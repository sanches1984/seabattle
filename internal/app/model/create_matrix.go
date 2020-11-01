package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	minRange = 4
	maxRange = 26
)

type CreateMatrixRequest struct {
	Range uint `json:"range"`
}

func NewCreateMatrixRequest(r *http.Request) (*CreateMatrixRequest, error) {
	var request CreateMatrixRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	if request.Range < minRange || request.Range > maxRange {
		return nil, fmt.Errorf("Range should be in range [%d:%d]", minRange, maxRange)
	}
	return &request, nil
}
