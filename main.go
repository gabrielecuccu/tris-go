package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"tris/ui"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tris")
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.SetFixedSize(true)

	state := ui.NewState()

	gameContainer := ui.NewGameContainer(state)

	myWindow.SetContent(gameContainer)
	myWindow.ShowAndRun()
}
