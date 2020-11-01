package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ShotRequest struct {
	Coord string `json:"coord"`
}

type Shot struct {
	X uint
	Y uint
}

type ShotResponse struct {
	Destroy bool `json:"destroy"`
	Knock   bool `json:"knock"`
	End     bool `json:"end"`
}

func NewShotRequest(r *http.Request) (*ShotRequest, error) {
	var request ShotRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	if request.Coord == "" {
		return nil, fmt.Errorf("Coord cannot be empty")
	}
	return &request, nil
}

func (r *ShotRequest) GetShot() *Shot {
	return &Shot{
		X: getNumber(strings.Join(numbers.FindAllString(r.Coord, -1), "")),
		Y: getNumberByLetter(strings.Join(letters.FindAllString(r.Coord, -1), "")),
	}
}

func (r *ShotResponse) GetJSON() []byte {
	bytes, _ := json.Marshal(r)
	return bytes
}

func (s *Shot) OverField(size uint) bool {
	return s.X >= size || s.Y >= size
}
