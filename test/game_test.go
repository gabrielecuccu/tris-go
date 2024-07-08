package test

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
	"tris/lib"
)

func TestNewGame(t *testing.T) {
	newGame := lib.NewGame()
	newGameStr := fmt.Sprintf("%s", newGame)
	wanted := fmt.Sprintf("[0 0 0]\n[0 0 0]\n[0 0 0]\n")
	assert.Equal(t, newGameStr, wanted)

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			assert.Equal(t, newGame.GameField[row][col], lib.EmptyCell)
		}
	}
}

func TestNewMove(t *testing.T) {
	aGame := lib.NewGame()
	aGame.NewMove(0, 2, lib.ComputerPlayer)
	aGame.NewMove(1, 0, lib.HumanPlayer)

	aGameStr := fmt.Sprintf("%s", aGame)
	wanted := fmt.Sprintf("[0 0 1]\n[2 0 0]\n[0 0 0]\n")
	assert.Equal(t, aGameStr, wanted)

	assert.Equal(t, aGame.GameField[0][0], lib.EmptyCell)
	assert.Equal(t, aGame.GameField[0][1], lib.EmptyCell)
	assert.Equal(t, aGame.GameField[0][2], lib.ComputerCell)
	assert.Equal(t, aGame.GameField[1][0], lib.HumanCell)
	assert.Equal(t, aGame.GameField[1][1], lib.EmptyCell)
	assert.Equal(t, aGame.GameField[1][2], lib.EmptyCell)
	assert.Equal(t, aGame.GameField[2][0], lib.EmptyCell)
	assert.Equal(t, aGame.GameField[2][1], lib.EmptyCell)
	assert.Equal(t, aGame.GameField[2][2], lib.EmptyCell)
}

func TestCanWin(t *testing.T) {
    // new game, nobody can win
	newGame := lib.NewGame()
	assert.Equal(t, newGame.CanWin(lib.ComputerPlayer), lib.Winner{false, 0, 0})
	assert.Equal(t, newGame.CanWin(lib.HumanPlayer), lib.Winner{false, 0, 0})

    // nobody can win
	newGame.GameField = &lib.Field{
		{lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
		{lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
		{lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
	}
	assert.Equal(t, newGame.CanWin(lib.ComputerPlayer), lib.Winner{false, 0, 0})
	assert.Equal(t, newGame.CanWin(lib.HumanPlayer), lib.Winner{false, 0, 0})

    // computer can win
    newGame.GameField = &lib.Field{
        {lib.ComputerCell, lib.HumanCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
    assert.Equal(t, newGame.CanWin(lib.ComputerPlayer), lib.Winner{true, 1, 0})
    assert.Equal(t, newGame.CanWin(lib.HumanPlayer), lib.Winner{false, 0, 0})

    // human can win
    newGame.GameField = &lib.Field{
        {lib.ComputerCell, lib.HumanCell, lib.HumanCell},
        {lib.HumanCell, lib.EmptyCell, lib.ComputerCell},
        {lib.ComputerCell, lib.HumanCell, lib.EmptyCell},
    }
    assert.Equal(t, newGame.CanWin(lib.ComputerPlayer), lib.Winner{false, 0, 0})
    assert.Equal(t, newGame.CanWin(lib.HumanPlayer), lib.Winner{true, 1, 1})

    // both human and computer can win
    newGame.GameField = &lib.Field{
        {lib.ComputerCell, lib.HumanCell, lib.ComputerCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
    assert.Equal(t, newGame.CanWin(lib.ComputerPlayer), lib.Winner{true, 1, 0})
    assert.Equal(t, newGame.CanWin(lib.HumanPlayer), lib.Winner{true, 2, 1})
}

func TestHasWin(t *testing.T) {
    // new game, nobody can win
	newGame := lib.NewGame()
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

    // computer can win (horizontal 1)
    newGame.GameField = &lib.Field{
        {lib.ComputerCell, lib.ComputerCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

	// computer can win (horizontal 2)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.ComputerCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

	// computer can win (horizontal 3)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.ComputerCell, lib.ComputerCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)
	
    // human can win (horizontal 1)
    newGame.GameField = &lib.Field{
        {lib.HumanCell, lib.HumanCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

	// human can win (horizontal 2)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.HumanCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

	// human can win (horizontal 3)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.HumanCell, lib.HumanCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

    // computer can win (vertical 1)
    newGame.GameField = &lib.Field{
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

    // computer can win (vertical 2)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

    // computer can win (vertical 3)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

    // human can win (vertical 1)
    newGame.GameField = &lib.Field{
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

    // human can win (vertical 2)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

    // human can win (vertical 3)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

    // computer can win (diagonal 1)
    newGame.GameField = &lib.Field{
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

	// computer can win (diagonal 2)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), false)

    // human can win (diagonal 1)
    newGame.GameField = &lib.Field{
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)

	// human can win (diagonal 2)
    newGame.GameField = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, newGame.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, newGame.HasWin(lib.HumanPlayer), true)
}