package ui

import (
    "time"
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"tris/lib"
)

func newCellButton(state *State, row, col int, buttons *[3][3]*widget.Button) *widget.Button {
    var button *widget.Button
    field := state.field
    button = widget.NewButton("", func() {
        if ct, _ := state.computerTurn.Get(); ct == true {
            return
        }

        state.computerTurn.Set(true)

        field.NewMove(row, col, lib.HumanPlayer)
        button.SetText("H")

        if field.HasWin(lib.HumanPlayer) {
            state.description.Set("You won!")
            return
        }

        if field.IsGameEnded() == true {
            state.description.Set("Nobody won!")
            return
        }

        time.Sleep(1 * time.Second)

        winner := field.CanWin(lib.ComputerPlayer)
        if winner.Exists {
            field.NewMove(winner.Row, winner.Col, lib.ComputerPlayer)
            buttons[winner.Row][winner.Col].SetText("C")
            state.description.Set("Computer won!")
            state.computerTurn.Set(false)
            return
        }

        winner = field.CanWin(lib.HumanPlayer)
        if winner.Exists {
            field.NewMove(winner.Row, winner.Col, lib.ComputerPlayer)
            buttons[winner.Row][winner.Col].SetText("C")

            if field.IsGameEnded() == true {
                state.description.Set("Nobody won!")
                return
            }

            state.description.Set("Your turn")
            state.computerTurn.Set(false)
            return
        }

        rRow, rCol := field.GetRandomEmptyCell()
        field.NewMove(rRow, rCol, lib.ComputerPlayer)
        buttons[rRow][rCol].SetText("C")

        if field.IsGameEnded() == true {
            state.description.Set("Nobody won!")
            return
        }

        state.description.Set("Your turn")
        state.computerTurn.Set(false)
    })

    buttons[row][col] = button

    return  button
}

func NewGameContainer(state *State) *fyne.Container {
    buttons := &[3][3]*widget.Button{}

    button1 := newCellButton(state, 0, 0, buttons)
    button2 := newCellButton(state, 0, 1, buttons)
    button3 := newCellButton(state, 0, 2, buttons)
    button4 := newCellButton(state, 1, 0, buttons)
    button5 := newCellButton(state, 1, 1, buttons)
    button6 := newCellButton(state, 1, 2, buttons)
    button7 := newCellButton(state, 2, 0, buttons)
    button8 := newCellButton(state, 2, 1, buttons)
    button9 := newCellButton(state, 2, 2, buttons)

    fieldContainer := container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9)

    statusLabel := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    statusLabel.Bind(state.description)

	contentContainer := container.NewBorder(nil, statusLabel, nil, nil, fieldContainer)

	return contentContainer
}