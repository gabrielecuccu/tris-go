package main

import (
	"fmt"
)

func main() {
	game := newGame()
	game.newMove(0, 0, computerPlayer)
	game.newMove(2, 2, computerPlayer)
	fmt.Println(game.canWin(computerPlayer))
	fmt.Println(game)
}
