package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"video-editor/app"
)

type VideoList struct {
	widget.List
	state *app.State
}

func NewVideoList(state *app.State) *VideoList {
	vl := &VideoList{state: state}
	vl.ExtendBaseWidget(vl)

	vl.Length = func() int {
		return state.Count()
	}

	vl.CreateItem = func() fyne.CanvasObject {
		img := canvas.NewImageFromImage(nil)
		img.SetMinSize(fyne.NewSize(120, 68))
		img.FillMode = canvas.ImageFillContain

		label := widget.NewLabel("placeholder")

		return container.NewHBox(img, label)
	}

	vl.UpdateItem = func(id widget.ListItemID, obj fyne.CanvasObject) {
		videos := state.GetVideos()
		if id < len(videos) {
			video := videos[id]
			box := obj.(*fyne.Container)
			img := box.Objects[0].(*canvas.Image)
			label := box.Objects[1].(*widget.Label)

			if video.Thumbnail != nil {
				img.Image = video.Thumbnail
				img.Refresh()
			}

			label.SetText(fmt.Sprintf("%d. %s (%s)", id+1, video.Name, video.SizeString()))
		}
	}

	vl.OnSelected = func(id widget.ListItemID) {
		state.SetSelected(int(id))
	}

	vl.OnUnselected = func(id widget.ListItemID) {
		state.SetSelected(-1)
	}

	return vl
}

func (vl *VideoList) Refresh() {
	vl.List.Refresh()

	selected := vl.state.GetSelected()
	if selected >= 0 && selected < vl.state.Count() {
		vl.Select(widget.ListItemID(selected))
	}
}
