package game

import (
	"fmt"
	"github.com/sanches1984/seabattle/internal/app/model"
)

// 0 - пустая ячейка
// 1 - корабль
// -1 - не попал
// -2 - подбитый корабль

type IClient interface {
	Clear()
	IsStarted() bool
	CreateShips(ships model.ShipList) error
	MakeShot(shot *model.Shot) (*model.ShotResponse, error)
	GetStat() *model.StateResponse
	GetInfo() string
}

type client struct {
	matrix    [][]int
	size      int
	started   bool
	ships     model.ShipList
	shotCount uint
}

func NewGame(size uint) IClient {
	c := &client{
		size:      int(size),
		started:   false,
		ships:     model.ShipList{},
		shotCount: 0,
	}
	c.matrix = c.createMatrix()
	return c
}

func (c *client) Clear() {
	c.started = true
	c.shotCount = 0
	c.ships.ResetHits()
	c.clearMatrix()
}

func (c client) IsStarted() bool {
	return c.started
}

func (c *client) CreateShips(ships model.ShipList) error {
	// создаем временное поле для проверки пересечения кораблей и выхода за пределы поля
	tempMatrix := c.createMatrix()
	for _, ship := range ships {
		ship.SetHits()
		if ship.OverField(uint(c.size)) {
			return fmt.Errorf("Ship out of field")
		}

		for i := ship.X1; i <= ship.X2; i++ {
			for j := ship.Y1; j <= ship.Y2; j++ {
				if tempMatrix[i][j] != 0 {
					return fmt.Errorf("Ships crossing")
				}
				tempMatrix[i][j] = 1
			}
		}
	}

	// корабли поставились нормально, ставим на настоящее поле
	c.matrix = tempMatrix
	c.ships = ships
	c.started = true
	return nil
}

func (c *client) MakeShot(shot *model.Shot) (*model.ShotResponse, error) {
	if shot.OverField(uint(c.size)) {
		return nil, fmt.Errorf("Shot out of field")
	}

	isHit := false
	switch c.matrix[shot.X][shot.Y] {
	case 0:
		c.matrix[shot.X][shot.Y] = -1
	case 1:
		c.matrix[shot.X][shot.Y] = -2
		isHit = true
	default:
		return nil, fmt.Errorf("Bad shot")
	}

	c.shotCount++
	response := &model.ShotResponse{}
	if isHit {
		response.Knock = true
		response.Destroy = c.ships.MakeDamage(shot)
		response.End = c.ships.AllShipsDestroyed()
		if response.End {
			c.started = false
		}
	}

	return response, nil
}

func (c client) GetStat() *model.StateResponse {
	return &model.StateResponse{
		ShipCount: len(c.ships),
		Destroyed: c.ships.Destroyed(),
		Knocked:   c.ships.Knocked(),
		ShotCount: int(c.shotCount),
	}
}

func (c client) GetInfo() string {
	str := "+--+"
	for i := 0; i < c.size; i++ {
		str += "-"
	}
	str += "+\n|  |"
	// печатаем первую строку
	for i := 0; i < c.size; i++ {
		// конвертируем в ASCII
		str += string(rune(i + 65))
	}
	str += "|\n+--+"
	for i := 0; i < c.size; i++ {
		str += "-"
	}
	str += "+\n"

	for i := 0; i < c.size; i++ {
		str += "|"
		if i+1 < 10 {
			str += "0"
		}
		str += fmt.Sprintf("%d|", i+1)
		for j := 0; j < c.size; j++ {
			switch c.matrix[i][j] {
			case 0:
				str += " "
			case 1:
				str += "O"
			case -1:
				str += "."
			case -2:
				str += "X"
			}
		}
		str += "|\n"
	}

	str += "+--+"
	for i := 0; i < c.size; i++ {
		str += "-"
	}
	str += "+\n"
	return str
}

func (c *client) createMatrix() [][]int {
	matrix := make([][]int, c.size)
	for i := 0; i < c.size; i++ {
		matrix[i] = make([]int, c.size)
	}
	return matrix
}

func (c *client) clearMatrix() {
	for i := 0; i < c.size; i++ {
		for j := 0; j < c.size; j++ {
			switch c.matrix[i][j] {
			case -1:
				c.matrix[i][j] = 0
			case -2:
				c.matrix[i][j] = 1
			}
		}
	}
}
