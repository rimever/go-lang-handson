package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type cdata struct {
	mem int
	cal string
	flg bool
}

func createNumButtons(f func(v int)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		createNumButton(f, 7),
		createNumButton(f, 8),
		createNumButton(f, 9),
		createNumButton(f, 4),
		createNumButton(f, 5),
		createNumButton(f, 6),
		createNumButton(f, 1),
		createNumButton(f, 2),
		createNumButton(f, 3),
	)
	return c
}

func createNumButton(f func(v int), i int) *widget.Button {
	return widget.NewButton(strconv.Itoa(i), func() { f(i) })
}
func createCalcButtons(f func(c string)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		createCalcButton(f, "CL"),
		createCalcButton(f, "/"),
		createCalcButton(f, "*"),
		createCalcButton(f, "+"),
		createCalcButton(f, "-"),
	)
	return c
}

func createCalcButton(f func(c string), s string) *widget.Button {
	return widget.NewButton(s, func() {
		f(s)
	})
}

func main() {
	a := app.New()
	w := a.NewWindow("Calc")
	w.SetFixedSize(true)
	l := widget.NewLabel("0")
	l.Alignment = fyne.TextAlignTrailing
	data := cdata{
		mem: 0,
		cal: "",
		flg: false,
	}
	calc := func(n int) {
		switch data.cal {
		case "":
			data.mem = n
		case "+":
			data.mem += n
		case "-":
			data.mem -= n
		case "*":
			data.mem *= n
		case "/":
			data.mem /= n
		}
		l.SetText(strconv.Itoa(data.mem))
		data.flg = true
	}
	pushNum := func(v int) {
		s := l.Text
		if data.flg {
			s = "0"
			data.flg = false
		}
		s += strconv.Itoa(v)
		n, err := strconv.Atoi(s)
		if err == nil {
			l.SetText(strconv.Itoa(n))
		}
	}
	pushCalc := func(c string) {
		if c == "CL" {
			l.SetText("0")
			data.mem = 0
			data.flg = false
			data.cal = ""
		}
		n, err := strconv.Atoi(l.Text)
		if err != nil {
			return
		}
		calc(n)
		data.cal = c
	}
	pushEnter := func() {
		n, err := strconv.Atoi(l.Text)
		if err != nil {
			return
		}
		calc(n)
		data.cal = ""
	}
	k := createNumButtons(pushNum)
	c := createCalcButtons(pushCalc)
	e := widget.NewButton("Enter", pushEnter)
	w.SetContent(
		fyne.NewContainerWithLayout(layout.NewBorderLayout(l, e, nil, c), l, e, k, c))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
