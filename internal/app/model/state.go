package model

import "encoding/json"

type StateResponse struct {
	ShipCount int `json:"ship_count"` // всего кораблей
	Destroyed int `json:"destroyed"`  // потоплено
	Knocked   int `json:"knocked"`    // подбито
	ShotCount int `json:"shot_count"` // сделано выстрелов
}

func (r *StateResponse) GetJSON() []byte {
	bytes, _ := json.Marshal(r)
	return bytes
}
