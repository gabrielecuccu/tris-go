package ui

import (
    "fmt"
    "time"
    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/gabrielecuccu/tris-go/lib"
)

func startNewGame(state *State, buttons *[3][3]*widget.Button, description string, computerTurn bool) {
    state.description.Set(description)
    state.computerTurn.Set(computerTurn)
    state.field.EmptyTheField()
    for row := 0; row < 3; row++ {
        for col := 0; col < 3; col++ {
            buttons[row][col].SetText("")
            buttons[row][col].SetIcon(nil)
        }
    }
}

func newHumanMove(state *State, row, col int, button *widget.Button, buttons *[3][3]*widget.Button) {
    field := state.field
    if ct, _ := state.computerTurn.Get(); ct == true {
        return
    }

    if state.field[row][col] != lib.EmptyCell {
        return
    }

    state.computerTurn.Set(true)

    field.NewMove(row, col, lib.HumanPlayer)
    button.SetIcon(GetHumanIcon())

    if field.HasWin(lib.HumanPlayer) {
        state.description.Set("You won!")
        state.incHumanScore()

        time.Sleep(2 * time.Second)

        // a new game starts here
        startNewGame(state, buttons, "Computer's turn", true)
    }

    if field.IsGameEnded() == true {
        state.description.Set("Nobody won!")
        time.Sleep(2 * time.Second)

        // a new game starts here
        startNewGame(state, buttons, "Computer's turn", true)
    }

    state.description.Set("Computer's turn")
    time.Sleep(1 * time.Second)

    winner := field.CanWin(lib.ComputerPlayer)
    if winner.Exists {
        field.NewMove(winner.Row, winner.Col, lib.ComputerPlayer)
        buttons[winner.Row][winner.Col].SetIcon(GetComputerIcon())
        state.description.Set("Computer won!")
        state.incComputerScore()
        time.Sleep(2 * time.Second)

        // a new game starts here
        startNewGame(state, buttons, "Your turn", false)
        return
    }

    winner = field.CanWin(lib.HumanPlayer)
    if winner.Exists {
        field.NewMove(winner.Row, winner.Col, lib.ComputerPlayer)
        buttons[winner.Row][winner.Col].SetIcon(GetComputerIcon())

        if field.IsGameEnded() == true {
            state.description.Set("Nobody won!")
            time.Sleep(2 * time.Second)

            // a new game starts here
            startNewGame(state, buttons, "Your turn", false)
        }

        state.description.Set("Your turn")
        state.computerTurn.Set(false)
        return
    }

    rRow, rCol := field.GetRandomEmptyCell()
    field.NewMove(rRow, rCol, lib.ComputerPlayer)
    buttons[rRow][rCol].SetIcon(GetComputerIcon())

    if field.IsGameEnded() == true {
        state.description.Set("Nobody won!")
        time.Sleep(2 * time.Second)

        // a new game starts here
        startNewGame(state, buttons, "Your turn", false)
    }

    state.description.Set("Your turn")
    state.computerTurn.Set(false)
}

func newCellButton(state *State, row, col int, buttons *[3][3]*widget.Button) *widget.Button {
    var button *widget.Button
    button = widget.NewButton("", func() {
        newHumanMove(state, row, col, button, buttons)
    })

    buttons[row][col] = button

    return  button
}

func newGameField(state *State) *fyne.Container {
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

    return container.New(layout.NewGridLayout(3), button1, button2, button3, button4, button5, button6, button7, button8, button9)
}

func newStatusBarTop(state *State) *fyne.Container {
    humanScoreLabel := widget.NewLabelWithStyle("Your score:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    humanScoreValue := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    state.scoreHuman.AddListener(binding.NewDataListener(func() {
        newScore, _ := state.scoreHuman.Get()
        humanScoreValue.SetText(fmt.Sprintf("%d", newScore))
    }))

    computerScoreLabel := widget.NewLabelWithStyle("Computer's score:", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    computerScoreValue := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    state.scoreComputer.AddListener(binding.NewDataListener(func() {
        newScore, _ := state.scoreComputer.Get()
        computerScoreValue.SetText(fmt.Sprintf("%d", newScore))
    }))

    return container.New(layout.NewGridLayout(4), humanScoreLabel, humanScoreValue, computerScoreLabel, computerScoreValue)
}

func NewGameContainer(state *State) *fyne.Container {
    gameField := newGameField(state)
    
    statusBarBottom := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
    statusBarBottom.Bind(state.description)

    statusBarTop := newStatusBarTop(state)

	contentContainer := container.NewBorder(statusBarTop, statusBarBottom, nil, nil, gameField)

	return contentContainer
}