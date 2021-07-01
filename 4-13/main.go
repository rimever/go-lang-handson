package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	tabs := container.NewAppTabs(
		container.NewTabItem("First", widget.NewLabel("This is First tab item.")),
		container.NewTabItem("Second",widget.NewLabel("This is Second tab item.")),
		)
	tabs.SetTabLocation(container.TabLocationTop)
	w.SetContent(tabs)
	w.ShowAndRun()
}