package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewEmptyFieldSize(t *testing.T) {
	emptyField := newEmptyField()
	assert.Equal(t, len(emptyField), 3)
	for row := 0; row < 3; row++ {
		assert.Equal(t, len(emptyField[row]), 3)
	}
}

func TestNewEmptyFieldContent(t *testing.T) {
	emptyField := newEmptyField()
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			assert.Equal(t, emptyField[row][col], emptyCell)
		}
	}
}
