package main

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewGame(t *testing.T) {
	newGame := newGame()
	newGameStr := fmt.Sprintf("%s", newGame)
	wanted := fmt.Sprintf("[0 0 0]\n[0 0 0]\n[0 0 0]\n")
	assert.Equal(t, newGameStr, wanted)

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			assert.Equal(t, newGame.gameField[row][col], emptyCell)
		}
	}
}

func TestNewMove(t *testing.T) {
	aGame := newGame()
	aGame.newMove(0, 2, computerPlayer)
	aGame.newMove(1, 0, humanPlayer)

	aGameStr := fmt.Sprintf("%s", aGame)
	wanted := fmt.Sprintf("[0 0 1]\n[2 0 0]\n[0 0 0]\n")
	assert.Equal(t, aGameStr, wanted)

	assert.Equal(t, aGame.gameField[0][0], emptyCell)
	assert.Equal(t, aGame.gameField[0][1], emptyCell)
	assert.Equal(t, aGame.gameField[0][2], computerCell)
	assert.Equal(t, aGame.gameField[1][0], humanCell)
	assert.Equal(t, aGame.gameField[1][1], emptyCell)
	assert.Equal(t, aGame.gameField[1][2], emptyCell)
	assert.Equal(t, aGame.gameField[2][0], emptyCell)
	assert.Equal(t, aGame.gameField[2][1], emptyCell)
	assert.Equal(t, aGame.gameField[2][2], emptyCell)
}
