package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"video-editor/app"
)

type MainLayout struct {
	Container *fyne.Container
	VideoList *VideoList
}

func NewMainLayout(state *app.State, handlers *app.Handlers) *MainLayout {
	videoList := NewVideoList(state)

	toolbar := NewToolbar(ToolbarHandlers{
		OnAdd:      handlers.OnAddVideos,
		OnRemove:   handlers.OnRemove,
		OnMoveUp:   handlers.OnMoveUp,
		OnMoveDown: handlers.OnMoveDown,
		OnClear:    handlers.OnClear,
		OnExport:   handlers.OnExport,
	})

	header := widget.NewLabel("Video Files (drag to reorder)")
	header.TextStyle = fyne.TextStyle{Bold: true}

	content := container.NewBorder(
		container.NewVBox(toolbar, header),
		nil, nil, nil,
		videoList,
	)

	return &MainLayout{
		Container: content,
		VideoList: videoList,
	}
}
