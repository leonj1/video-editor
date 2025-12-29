package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ToolbarHandlers struct {
	OnAdd      func()
	OnRemove   func()
	OnMoveUp   func()
	OnMoveDown func()
	OnClear    func()
}

func NewToolbar(handlers ToolbarHandlers) fyne.CanvasObject {
	addBtn := widget.NewButtonWithIcon("Add Videos", theme.ContentAddIcon(), handlers.OnAdd)
	removeBtn := widget.NewButtonWithIcon("Remove", theme.ContentRemoveIcon(), handlers.OnRemove)
	upBtn := widget.NewButtonWithIcon("Move Up", theme.MoveUpIcon(), handlers.OnMoveUp)
	downBtn := widget.NewButtonWithIcon("Move Down", theme.MoveDownIcon(), handlers.OnMoveDown)
	clearBtn := widget.NewButtonWithIcon("Clear All", theme.DeleteIcon(), handlers.OnClear)

	return container.NewHBox(
		addBtn,
		removeBtn,
		widget.NewSeparator(),
		upBtn,
		downBtn,
		widget.NewSeparator(),
		clearBtn,
	)
}
