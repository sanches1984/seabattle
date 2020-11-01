package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type ShipList []*Ship

var numbers = regexp.MustCompile("[0-9]+")
var letters = regexp.MustCompile("[A-Z]+")

type ShipRequest struct {
	Coordinates string `json:"Coordinates"`
}

type Ship struct {
	X1          uint
	Y1          uint
	X2          uint
	Y2          uint
	TotalHits   uint
	CurrentHits uint
}

func NewShipRequest(r *http.Request) (*ShipRequest, error) {
	var request ShipRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	if request.Coordinates == "" {
		return nil, fmt.Errorf("Coordinates cannot be empty")
	}
	return &request, nil
}

func (r *ShipRequest) GetShips() (ShipList, error) {
	coords := strings.Split(r.Coordinates, ",")

	ships := make([]*Ship, 0, len(coords))
	for _, coord := range coords {
		xy := strings.Split(coord, " ")
		if len(xy) != 2 {
			return nil, fmt.Errorf("Bad coordinates")
		}

		ship := &Ship{
			X1: getNumber(strings.Join(numbers.FindAllString(xy[0], -1), "")),
			Y1: getNumberByLetter(strings.Join(letters.FindAllString(xy[0], -1), "")),
			X2: getNumber(strings.Join(numbers.FindAllString(xy[1], -1), "")),
			Y2: getNumberByLetter(strings.Join(letters.FindAllString(xy[1], -1), "")),
		}
		ship.SetHits()
		ships = append(ships, ship)
	}
	return ships, nil
}

func (list ShipList) MakeDamage(shot *Shot) bool {
	for _, ship := range list {
		if shot.X >= ship.X1 && shot.X <= ship.X2 && shot.Y >= ship.Y1 && shot.Y <= ship.Y2 {
			ship.CurrentHits--
			if ship.CurrentHits == 0 {
				// убит
				return true
			}
			return false
		}
	}
	return false
}

func (list ShipList) ResetHits() {
	for _, ship := range list {
		ship.CurrentHits = ship.TotalHits
	}
}

func (list ShipList) AllShipsDestroyed() bool {
	for _, ship := range list {
		if ship.CurrentHits > 0 {
			return false
		}
	}
	return true
}

func (list ShipList) Destroyed() int {
	cnt := 0
	for _, ship := range list {
		if ship.CurrentHits == 0 {
			cnt++
		}
	}
	return cnt
}

func (list ShipList) Knocked() int {
	cnt := 0
	for _, ship := range list {
		if ship.CurrentHits != 0 && ship.CurrentHits != ship.TotalHits {
			cnt++
		}
	}
	return cnt
}

func (s *Ship) OverField(size uint) bool {
	return s.X1 >= size || s.X2 >= size || s.Y1 >= size || s.Y2 >= size
}

func (s *Ship) SetHits() {
	s.TotalHits = (s.Y2 - s.Y1 + 1) * (s.X2 - s.X1 + 1)
	s.CurrentHits = s.TotalHits
}

func getNumberByLetter(l string) uint {
	// делаем сдвиг по ASCII
	return uint(([]rune(l))[0]) - 65
}

func getNumber(l string) uint {
	// ошибок нет будет, вырезли только цифры
	v, _ := strconv.Atoi(l)
	// приводим к счету с 0
	return uint(v - 1)
}
