package main

import (
    "fmt"
)

type cell int

type field [3][3]cell

type position int

type player int

const (
    ONE   position = 0
    TWO   position = 1
    THREE position = 2
)

const (
	emptyCell    cell = 0
	computerCell cell = 1
	humanCell    cell = 2
)

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

type game struct {
    gameField *field
}

func (g game) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n", g.gameField[0], g.gameField[1], g.gameField[2])
}

func (g game) newMove(row, col position, p player) {
    if cell := g.gameField[row][col]; cell == emptyCell {
        g.gameField[row][col] = p.toCell()
    }
}

func (g game) canComputerWin() bool {
    return false
}

func (g game) hasWin(p player) bool {
    c := p.toCell()
    f := g.gameField
    for row := 0; row < 3; row++ {
        if f[row][0] == c && f[row][1] == c && f[row][2] == c {
            return true
        }
    }

    for col := 0; col < 3; col++ {
        if f[0][col] == c && f[1][col] == c && f[2][col] == c {
            return true
        }
    }
    
    if f[ONE][ONE] == c && f[TWO][TWO] == c && f[THREE][THREE] == c {
        return true
    }
    
    if f[ONE][THREE] == c && f[TWO][TWO] == c && f[THREE][ONE] == c {
        return true
    }

    return false
}

func newEmptyField() *field {
    return &field {
        {emptyCell, emptyCell, emptyCell},
        {emptyCell, emptyCell, emptyCell},
        {emptyCell, emptyCell, emptyCell},
    }
}

func newGame() *game {
    return &game {
        gameField: newEmptyField(),
    }
}

func main() {
    game := newGame()
    game.newMove(ONE, THREE, computerPlayer)
    game.newMove(TWO, THREE, computerPlayer)
    game.newMove(THREE, THREE, computerPlayer)
    fmt.Println(game)
    fmt.Println(game.hasWin(computerPlayer))
}
