package main

import (
    "fmt"
)

func main() {
    game := newGame()
    game.newMove(ONE, THREE, computerPlayer)
    game.newMove(TWO, THREE, computerPlayer)
    game.newMove(THREE, THREE, computerPlayer)
    fmt.Println(game)
    fmt.Println(game.hasWin(computerPlayer))
}
