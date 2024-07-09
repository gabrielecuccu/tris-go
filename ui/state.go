package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"tris/lib"
)

type State struct {
    Description binding.String
    scoreHuman binding.Int
    scoreComputer binding.Int
    field *lib.Field
}

func NewState() *State {
    state := &State {
        Description: binding.NewString(),
        scoreHuman: binding.NewInt(),
        scoreComputer: binding.NewInt(),
        field: lib.NewEmptyField(),
    }

    state.Description.Set("You: 0 - Computer: 0 | Your turn")
    state.scoreComputer.Set(0)
    state.scoreHuman.Set(0)

    return state
}