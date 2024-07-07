package main

import (
    "fmt"
)

func main() {
    game := newGame()
    game.newMove(ONE, ONE, computerPlayer)
//     game.newMove(TWO, TWO, computerPlayer)
    game.newMove(THREE, THREE, computerPlayer)
    fmt.Println(game.canComputerWin())
    fmt.Println(game)
}
