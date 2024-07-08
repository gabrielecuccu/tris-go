package main

type player int

const (
	computerPlayer player = 1
	humanPlayer    player = 2
)

func (p player) toCell() cell {
	switch p {
	case computerPlayer:
		return computerCell
	case humanPlayer:
		return humanCell
	default:
		return emptyCell
	}
}
