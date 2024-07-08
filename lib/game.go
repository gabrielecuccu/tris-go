package lib

import (
	"fmt"
)

type game struct {
	GameField *Field
}

type Winner struct {
	Exists bool
	Row    int
	Col    int
}

func (g game) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n", g.GameField [0], g.GameField [1], g.GameField [2])
}

func (g game) NewMove(row, col int, p Player) {
	if cell := g.GameField [row][col]; cell == EmptyCell {
		g.GameField [row][col] = p.ToCell()
	}
}

func (g game) CanWin(p Player) Winner {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if curCell := g.GameField [row][col]; curCell == EmptyCell {
				g.GameField [row][col] = p.ToCell()
				canWin := g.hasWin(p)
				g.GameField [row][col] = EmptyCell
				if canWin == true {
					return Winner{true, row, col}
				}
			}
		}
	}

	return Winner{false, 0, 0}
}

func (g game) hasWin(p Player) bool {
	c := p.ToCell()
	f := g.GameField 
	for row := 0; row < 3; row++ {
		if f[row][0] == c && f[row][1] == c && f[row][2] == c {
			return true
		}
	}

	for col := 0; col < 3; col++ {
		if f[0][col] == c && f[1][col] == c && f[2][col] == c {
			return true
		}
	}

	if f[0][0] == c && f[1][1] == c && f[2][2] == c {
		return true
	}

	if f[0][2] == c && f[1][1] == c && f[2][0] == c {
		return true
	}

	return false
}

func NewGame() *game {
	return &game{
		GameField: NewEmptyField(),
	}
}
