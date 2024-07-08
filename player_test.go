package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestComputerPlayerToCell(t *testing.T) {
	assert.Equal(t, computerPlayer.toCell(), computerCell)
	assert.Equal(t, humanPlayer.toCell(), humanCell)
	assert.Equal(t, player(6).toCell(), emptyCell)
}
