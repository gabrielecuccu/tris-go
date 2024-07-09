package ui

import (
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
)

func newCellButton() *widget.Button {
    cellButton := widget.NewButton("", func() {

    })

    return  cellButton
}

func NewGameContainer(state *State) *fyne.Container {
    button1 := newCellButton()
    button2 := newCellButton()
    button3 := newCellButton()
    button4 := newCellButton()
    button5 := newCellButton()
    button6 := newCellButton()
    button7 := newCellButton()
    button8 := newCellButton()
    button9 := newCellButton()

    fieldContainer := container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9)

    statusLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    statusLabel.Bind(state.Description)

	contentContainer := container.NewBorder(nil, statusLabel, nil, nil, fieldContainer)

	return contentContainer
}