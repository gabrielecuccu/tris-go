package test

import (
	"gotest.tools/v3/assert"
	"testing"
	"tris/lib"
)

func TestNewEmptyFieldSize(t *testing.T) {
	emptyField := lib.NewEmptyField()
	assert.Equal(t, len(emptyField), 3)
	for row := 0; row < 3; row++ {
		assert.Equal(t, len(emptyField[row]), 3)
	}
}

func TestNewEmptyFieldContent(t *testing.T) {
	emptyField := lib.NewEmptyField()
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			assert.Equal(t, emptyField[row][col], lib.EmptyCell)
		}
	}
}
