package ui

import (
	"fyne.io/fyne/v2/data/binding"
)

type State struct {
    Description binding.String
}

func NewState() *State {
    state := &State {
        Description: binding.NewString(),
    }

    state.Description.Set("Welcome to Tris")

    return state
}