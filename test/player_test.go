package test

import (
	"gotest.tools/v3/assert"
	"testing"
	"tris/lib"
)

func TestComputerPlayerToCell(t *testing.T) {
	assert.Equal(t, lib.ComputerPlayer.ToCell(), lib.ComputerCell)
	assert.Equal(t, lib.HumanPlayer.ToCell(), lib.HumanCell)
	assert.Equal(t, lib.Player(6).ToCell(), lib.EmptyCell)
}
