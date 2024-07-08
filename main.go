package main

import (
	"fmt"
	"tris/lib"
)

func main() {
	game := lib.NewGame()
	game.NewMove(0, 0, lib.ComputerPlayer)
	game.NewMove(2, 2, lib.ComputerPlayer)
	fmt.Println(game.CanWin(lib.ComputerPlayer))
	fmt.Println(game)
}
