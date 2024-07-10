package test

import (
	"gotest.tools/v3/assert"
	"testing"
	"tris/lib"
	"fmt"
)

func TestNewEmptyFieldSize(t *testing.T) {
	emptyField := lib.NewEmptyField()
	assert.Equal(t, len(emptyField), 3)
	for row := 0; row < 3; row++ {
		assert.Equal(t, len(emptyField[row]), 3)
	}
}

func TestNewEmptyFieldContent(t *testing.T) {
	emptyField := lib.NewEmptyField()
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			assert.Equal(t, emptyField[row][col], lib.EmptyCell)
		}
	}
}

func TestNewMove(t *testing.T) {
	field := lib.NewEmptyField()
	field.NewMove(0, 2, lib.ComputerPlayer)
	field.NewMove(1, 0, lib.HumanPlayer)

	fieldStr := fmt.Sprintf("%s", field)
	wanted := fmt.Sprintf("[0 0 1]\n[2 0 0]\n[0 0 0]\n")
	assert.Equal(t, fieldStr, wanted)

	assert.Equal(t, field[0][0], lib.EmptyCell)
	assert.Equal(t, field[0][1], lib.EmptyCell)
	assert.Equal(t, field[0][2], lib.ComputerCell)
	assert.Equal(t, field[1][0], lib.HumanCell)
	assert.Equal(t, field[1][1], lib.EmptyCell)
	assert.Equal(t, field[1][2], lib.EmptyCell)
	assert.Equal(t, field[2][0], lib.EmptyCell)
	assert.Equal(t, field[2][1], lib.EmptyCell)
	assert.Equal(t, field[2][2], lib.EmptyCell)
}

func TestCanWin(t *testing.T) {
    // new game, nobody can win
	field := lib.NewEmptyField()
	assert.Equal(t, field.CanWin(lib.ComputerPlayer), lib.Winner{false, 0, 0})
	assert.Equal(t, field.CanWin(lib.HumanPlayer), lib.Winner{false, 0, 0})

    // nobody can win
	field = &lib.Field{
		{lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
		{lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
		{lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
	}
	assert.Equal(t, field.CanWin(lib.ComputerPlayer), lib.Winner{false, 0, 0})
	assert.Equal(t, field.CanWin(lib.HumanPlayer), lib.Winner{false, 0, 0})

    // computer can win
    field = &lib.Field{
        {lib.ComputerCell, lib.HumanCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
    assert.Equal(t, field.CanWin(lib.ComputerPlayer), lib.Winner{true, 1, 0})
    assert.Equal(t, field.CanWin(lib.HumanPlayer), lib.Winner{false, 0, 0})

    // human can win
    field = &lib.Field{
        {lib.ComputerCell, lib.HumanCell, lib.HumanCell},
        {lib.HumanCell, lib.EmptyCell, lib.ComputerCell},
        {lib.ComputerCell, lib.HumanCell, lib.EmptyCell},
    }
    assert.Equal(t, field.CanWin(lib.ComputerPlayer), lib.Winner{false, 0, 0})
    assert.Equal(t, field.CanWin(lib.HumanPlayer), lib.Winner{true, 1, 1})

    // both human and computer can win
    field = &lib.Field{
        {lib.ComputerCell, lib.HumanCell, lib.ComputerCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
    assert.Equal(t, field.CanWin(lib.ComputerPlayer), lib.Winner{true, 1, 0})
    assert.Equal(t, field.CanWin(lib.HumanPlayer), lib.Winner{true, 2, 1})
}

func TestHasWin(t *testing.T) {
    // new game, nobody can win
	field := lib.NewEmptyField()
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

    // computer can win (horizontal 1)
    field = &lib.Field{
        {lib.ComputerCell, lib.ComputerCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

	// computer can win (horizontal 2)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.ComputerCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

	// computer can win (horizontal 3)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.ComputerCell, lib.ComputerCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)
	
    // human can win (horizontal 1)
    field = &lib.Field{
        {lib.HumanCell, lib.HumanCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

	// human can win (horizontal 2)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.HumanCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

	// human can win (horizontal 3)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.HumanCell, lib.HumanCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

    // computer can win (vertical 1)
    field = &lib.Field{
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

    // computer can win (vertical 2)
    field = &lib.Field{
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

    // computer can win (vertical 3)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

    // human can win (vertical 1)
    field = &lib.Field{
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

    // human can win (vertical 2)
    field = &lib.Field{
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

    // human can win (vertical 3)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

    // computer can win (diagonal 1)
    field = &lib.Field{
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

	// computer can win (diagonal 2)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.ComputerCell},
        {lib.EmptyCell, lib.ComputerCell, lib.EmptyCell},
        {lib.ComputerCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), true)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), false)

    // human can win (diagonal 1)
    field = &lib.Field{
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)

	// human can win (diagonal 2)
    field = &lib.Field{
        {lib.EmptyCell, lib.EmptyCell, lib.HumanCell},
        {lib.EmptyCell, lib.HumanCell, lib.EmptyCell},
        {lib.HumanCell, lib.EmptyCell, lib.EmptyCell},
    }
	assert.Equal(t, field.HasWin(lib.ComputerPlayer), false)
	assert.Equal(t, field.HasWin(lib.HumanPlayer), true)
}

func TestIsGameEnded(t *testing.T) {
	field := lib.NewEmptyField()
	assert.Equal(t, field.IsGameEnded(), false)

    field = &lib.Field{
        {lib.EmptyCell, lib.HumanCell, lib.ComputerCell},
        {lib.ComputerCell, lib.HumanCell, lib.HumanCell},
        {lib.ComputerCell, lib.EmptyCell, lib.ComputerCell},
    }
    assert.Equal(t, field.IsGameEnded(), false)

    field = &lib.Field{
        {lib.HumanCell, lib.HumanCell, lib.ComputerCell},
        {lib.ComputerCell, lib.ComputerCell, lib.HumanCell},
        {lib.HumanCell, lib.HumanCell, lib.ComputerCell},
    }
    assert.Equal(t, field.IsGameEnded(), true)
}

func TestEmptyTheField(t *testing.T) {
    field := lib.NewEmptyField()
    field = &lib.Field{
        {lib.HumanCell, lib.HumanCell, lib.ComputerCell},
        {lib.ComputerCell, lib.ComputerCell, lib.HumanCell},
        {lib.HumanCell, lib.HumanCell, lib.ComputerCell},
    }

    field.EmptyTheField()
    assert.Equal(t, *field, *lib.NewEmptyField())
}
