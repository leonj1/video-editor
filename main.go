package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	appPkg "video-editor/app"
	"video-editor/ui"
)

func main() {
	a := app.NewWithID("com.videoarranger.app")
	window := a.NewWindow("Video Arranger")

	state := appPkg.NewState()
	handlers := appPkg.NewHandlers(state, window)
	layout := ui.NewMainLayout(state, handlers)

	state.SetOnChange(func() {
		layout.VideoList.Refresh()
	})

	window.SetContent(layout.Container)
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
