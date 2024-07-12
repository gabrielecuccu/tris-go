package test

import (
    "github.com/gabrielecuccu/tris-go/lib"
	"gotest.tools/v3/assert"
	"testing"
)

func TestComputerPlayerToCell(t *testing.T) {
	assert.Equal(t, lib.ComputerPlayer.ToCell(), lib.ComputerCell)
	assert.Equal(t, lib.HumanPlayer.ToCell(), lib.HumanCell)
	assert.Equal(t, lib.Player(6).ToCell(), lib.EmptyCell)
}
