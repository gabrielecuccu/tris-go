package lib

import (
    "fmt"
    "math/rand/v2"
)

type Field [3][3]cell

type Winner struct {
	Exists bool
	Row    int
	Col    int
}

func NewEmptyField() *Field {
	return &Field{
		{EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell},
	}
}

func (field *Field) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n", field [0], field[1], field [2])
}

func (field *Field) NewMove(row, col int, p Player) {
	if cell := field [row][col]; cell == EmptyCell {
		field [row][col] = p.ToCell()
	}
}

func (field *Field) CanWin(p Player) Winner {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if curCell := field [row][col]; curCell == EmptyCell {
				field [row][col] = p.ToCell()
				canWin := field.HasWin(p)
				field [row][col] = EmptyCell
				if canWin == true {
					return Winner{true, row, col}
				}
			}
		}
	}

	return Winner{false, 0, 0}
}

func (field *Field) HasWin(p Player) bool {
	c := p.ToCell()
	for row := 0; row < 3; row++ {
		if field[row][0] == c && field[row][1] == c && field[row][2] == c {
			return true
		}
	}

	for col := 0; col < 3; col++ {
		if field[0][col] == c && field[1][col] == c && field[2][col] == c {
			return true
		}
	}

	if field[0][0] == c && field[1][1] == c && field[2][2] == c {
		return true
	}

	if field[0][2] == c && field[1][1] == c && field[2][0] == c {
		return true
	}

	return false
}

func (field *Field) GetRandomEmptyCell() (int, int) {
    for {
        row := rand.IntN(3)
        col := rand.IntN(3)
        if c := field[row][col]; c == EmptyCell {
            return row, col
        }
    }
}

func (field *Field) IsGameEnded() bool {
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            if field[row][col] == EmptyCell {
                return false
            }
        }
    }

    return true
}