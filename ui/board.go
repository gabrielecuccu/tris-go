package ui

import (
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"tris/lib"
)

func newCellButton(field *lib.Field, row, col int) *widget.Button {
    var button *widget.Button
    button = widget.NewButton("", func() {
        field.NewMove(row, col, lib.HumanPlayer)
        button.SetText("H")
    })

    return  button
}

func NewGameContainer(state *State) *fyne.Container {
    button1 := newCellButton(state.field, 0, 0)
    button2 := newCellButton(state.field, 0, 1)
    button3 := newCellButton(state.field, 0, 2)
    button4 := newCellButton(state.field, 1, 0)
    button5 := newCellButton(state.field, 1, 1)
    button6 := newCellButton(state.field, 1, 2)
    button7 := newCellButton(state.field, 2, 0)
    button8 := newCellButton(state.field, 2, 1)
    button9 := newCellButton(state.field, 2, 2)

    fieldContainer := container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9)

    statusLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    statusLabel.Bind(state.Description)

	contentContainer := container.NewBorder(nil, statusLabel, nil, nil, fieldContainer)

	return contentContainer
}