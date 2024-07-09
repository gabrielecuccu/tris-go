package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"tris/lib"
)

type State struct {
    description binding.String
    scoreHuman binding.Int
    scoreComputer binding.Int
    field *lib.Field
    computerTurn binding.Bool
}

func NewState() *State {
    state := &State {
        description: binding.NewString(),
        scoreHuman: binding.NewInt(),
        scoreComputer: binding.NewInt(),
        field: lib.NewEmptyField(),
        computerTurn: binding.NewBool(),
    }

    state.description.Set("Your turn")
    state.scoreComputer.Set(0)
    state.scoreHuman.Set(0)
    state.computerTurn.Set(false)

    return state
}