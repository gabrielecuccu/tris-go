package main

import (
	"testing"
	"fmt"
)

func TestNewGame(t *testing.T) {
    newGame := newGame()
    newGameStr := fmt.Sprintf("%s", newGame)
    wanted := fmt.Sprintf("[0 0 0]\n[0 0 0]\n[0 0 0]\n")
    if newGameStr != wanted {
        t.Errorf("wanted %s, got %s", wanted, newGameStr)
    }

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if c := newGame.gameField[row][col]; c != emptyCell {
				t.Errorf("cell [%d][%d] must be emptyCell, got %d", row, col, c)
			}
		}
	}
}

func TestNewMove(t *testing.T) {
    aGame := newGame()
    aGame.newMove(ONE, THREE, computerPlayer)
    aGame.newMove(TWO, ONE, humanPlayer)
    aGameStr := fmt.Sprintf("%s", aGame)

    wanted := fmt.Sprintf("[0 0 1]\n[2 0 0]\n[0 0 0]\n")
    if aGameStr != wanted {
        t.Errorf("wanted %s, got %s", wanted, aGameStr)
    }

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
		    c := aGame.gameField[row][col]
            if row == int(ONE) && col == int(THREE) {
                if c != computerCell {
                    t.Errorf("cell [1][3] must be computerCell, got %d", c)
                }
            } else if row == int(TWO) && col == int(ONE) {
                if c != humanCell {
                    t.Errorf("cell [2][1] must be humanCell, got %d", c)
                }
		    } else if c != emptyCell {
				t.Errorf("cell [%d][%d] must be emptyCell, got %d", row, col, c)
			}
		}
	}
}

