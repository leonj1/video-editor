package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ShowError(err error, window fyne.Window) {
	dialog.ShowError(err, window)
}

func ShowInfo(title, message string, window fyne.Window) {
	dialog.ShowInformation(title, message, window)
}
