package main

import (
	"fmt"
)

type game struct {
	gameField *field
}

type winner struct {
	exists bool
	row    position
	col    position
}

func (g game) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n", g.gameField[0], g.gameField[1], g.gameField[2])
}

func (g game) newMove(row, col position, p player) {
	if cell := g.gameField[row][col]; cell == emptyCell {
		g.gameField[row][col] = p.toCell()
	}
}

func (g game) canComputerWin() winner {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if curCell := g.gameField[row][col]; curCell == emptyCell {
				g.gameField[row][col] = computerCell
				canWin := g.hasWin(computerPlayer)
				g.gameField[row][col] = emptyCell
				if canWin == true {
					return winner{true, position(row), position(col)}
				}
			}
		}
	}

	return winner{false, 0, 0}
}

func (g game) hasWin(p player) bool {
	c := p.toCell()
	f := g.gameField
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

	if f[ONE][ONE] == c && f[TWO][TWO] == c && f[THREE][THREE] == c {
		return true
	}

	if f[ONE][THREE] == c && f[TWO][TWO] == c && f[THREE][ONE] == c {
		return true
	}

	return false
}

func newGame() *game {
	return &game{
		gameField: newEmptyField(),
	}
}
