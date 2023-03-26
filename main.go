package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

type Data struct {
	w fyne.Window
	p *widget.ProgressBar
}

var InDirectory string
var OutDirectory string

func main() {
	a := app.New()
	w := a.NewWindow("seren-stems")
	p := widget.NewProgressBar()

	d := Data{
		w: w,
		p: p,
	}

	renderMenu(d)

	w.ShowAndRun()
}

func process(d Data) {
	fmt.Println(InDirectory, OutDirectory)
}

func renderMenu(d Data) {
	inDirSelected := widget.NewLabel(InDirectory)
	outDirSelected := widget.NewLabel(OutDirectory)

	d.w.SetContent(
		container.NewVBox(
			container.NewHBox(
				widget.NewLabel("Input directory"),
				inDirSelected,
				widget.NewButton("select", func() {
					dir, err := dialog.Directory().Title("Select input directory").Browse()
					if err != nil {
						return
					}
					InDirectory = dir
					inDirSelected.SetText(dir)
				}),
			),
			container.NewHBox(
				widget.NewLabel("Output directory"),
				outDirSelected,
				widget.NewButton("select", func() {
					dir, err := dialog.Directory().Title("Select output directory").Browse()
					if err != nil {
						return
					}
					OutDirectory = dir
					outDirSelected.SetText(OutDirectory)
				}),
			),
			container.NewHBox(
				widget.NewButton("Start", func() {
					renderProcessing(d)
					process(d)
				}),
			),
		),
	)
}

func renderProcessing(d Data) {
	d.w.SetContent(
		container.NewVBox(
			container.NewHBox(
				widget.NewLabel("Processing..."),
			),
			container.NewHBox(
				d.p,
			),
			container.NewHBox(
				widget.NewButton("Cancel", func() {
					d.p.SetValue(0)
					renderMenu(d)
					// going to need to cancel the process too
					// probably want to have a cancel channel
					// and have the process check it
					// also want to clean up the temp files & report progress
				}),
			),
		),
	)
}
