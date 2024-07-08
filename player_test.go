package main

import (
	"testing"
)

func TestComputerPlayerToCell(t *testing.T) {
	p := computerPlayer
	c := p.toCell()
	if c != computerCell {
		t.Errorf("computerPlayer.toCell = %d; want %d", c, computerCell)
	}
}

func TestHumanPlayerToCell(t *testing.T) {
	p := humanPlayer
	c := p.toCell()
	if c != humanCell {
		t.Errorf("humanPlayer.toCell = %d; want %d", c, humanCell)
	}
}

func TestUnknownPlayerToCell(t *testing.T) {
	p := player(6)
	c := p.toCell()
	if c != emptyCell {
		t.Errorf("unknownPlayer.toCell = %d; want %d", c, emptyCell)
	}
}
