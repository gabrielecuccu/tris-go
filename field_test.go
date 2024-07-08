package main

import (
	"testing"
)

func TestNewEmptyFieldSize(t *testing.T) {
	emptyField := newEmptyField()
	if length := len(emptyField); length != 3 {
		t.Errorf("field length must have 3 rows, got %d", length)
	}

	for row := 0; row < 3; row++ {
		if length := len(emptyField[row]); length != 3 {
			t.Errorf("field %d rows must have 3 cols, got %d", row, length)
		}
	}
}

func TestNewEmptyFieldContent(t *testing.T) {
	emptyField := newEmptyField()
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if c := emptyField[row][col]; c != emptyCell {
				t.Errorf("cell [%d][%d] must be emptyCell, got %d", row, col, c)
			}
		}
	}
}
