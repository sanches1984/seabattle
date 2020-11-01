package game

import (
	"fmt"
	"github.com/sanches1984/seabattle/internal/app/model"
)

// 0 - пустая ячейка
// 1 - корабль
// -1 - не попал
// -2 - подбитый корабль
var matrix [][]int
var matrixSize int
var isGameStarted = false
var shipsInfo model.ShipList
var shotCount uint

func IsGameStarted() bool {
	return isGameStarted
}

func NewGame(size uint) {
	isGameStarted = false
	shotCount = 0
	shipsInfo = model.ShipList{}
	matrixSize = int(size)
	matrix = createMatrix()
}

func ClearGame() {
	isGameStarted = true
	shotCount = 0
	shipsInfo.ResetHits()
	clearMatrix()
}

func CreateShips(ships model.ShipList) error {
	// создаем временное поле для проверки пересечения кораблей и выхода за пределы поля
	tempMatrix := createMatrix()
	for _, ship := range ships {
		ship.SetHits()
		if ship.OverField(uint(matrixSize)) {
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
	matrix = tempMatrix
	shipsInfo = ships
	isGameStarted = true
	return nil
}

func MakeShot(shot *model.Shot) (*model.ShotResponse, error) {
	if shot.OverField(uint(matrixSize)) {
		return nil, fmt.Errorf("Shot out of field")
	}

	isHit := false
	switch matrix[shot.X][shot.Y] {
	case 0:
		matrix[shot.X][shot.Y] = -1
	case 1:
		matrix[shot.X][shot.Y] = -2
		isHit = true
	default:
		return nil, fmt.Errorf("Bad shot")
	}

	shotCount++
	response := &model.ShotResponse{}
	if isHit {
		response.Knock = true
		response.Destroy = shipsInfo.MakeDamage(shot)
		response.End = shipsInfo.AllShipsDestroyed()
		if response.End {
			isGameStarted = false
		}
	}

	return response, nil
}

func GetStat() *model.StateResponse {
	return &model.StateResponse{
		ShipCount: len(shipsInfo),
		Destroyed: shipsInfo.Destroyed(),
		Knocked:   shipsInfo.Knocked(),
		ShotCount: int(shotCount),
	}
}

func GetInfo() string {
	str := "+--+"
	for i := 0; i < matrixSize; i++ {
		str += "-"
	}
	str += "+\n|  |"
	// печатаем первую строку
	for i := 0; i < matrixSize; i++ {
		// конвертируем в ASCII
		str += string(rune(i + 65))
	}
	str += "|\n+--+"
	for i := 0; i < matrixSize; i++ {
		str += "-"
	}
	str += "+\n"

	for i := 0; i < matrixSize; i++ {
		str += "|"
		if i+1 < 10 {
			str += "0"
		}
		str += fmt.Sprintf("%d|", i+1)
		for j := 0; j < matrixSize; j++ {
			switch matrix[i][j] {
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
	for i := 0; i < matrixSize; i++ {
		str += "-"
	}
	str += "+\n"
	return str
}

func createMatrix() [][]int {
	matrix := make([][]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]int, matrixSize)
	}
	return matrix
}

func clearMatrix() {
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			switch matrix[i][j] {
			case -1:
				matrix[i][j] = 0
			case -2:
				matrix[i][j] = 1
			}
		}
	}
}
