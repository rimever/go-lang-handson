package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	c := 0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	w.SetContent(
		container.NewVBox(
			l,
			widget.NewButton("Click me!", func() {
				c++
				l.SetText("count: " + strconv.Itoa(c))
			})))
	w.ShowAndRun()
}