package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// CreateMainContent It creates the main content to display after initial startup
func notInUse(app fyne.App) *widget.Box {
	return widget.NewVBox(
		widget.NewLabel("Hello World!"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	)
}
