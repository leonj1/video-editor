package app

import (
	"log"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var videoExtensions = []string{".mp4", ".mov", ".avi", ".mkv", ".webm", ".m4v", ".wmv", ".flv"}

func isVideoFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	for _, e := range videoExtensions {
		if ext == e {
			return true
		}
	}
	return false
}

type Handlers struct {
	state  *State
	window fyne.Window
}

func NewHandlers(state *State, window fyne.Window) *Handlers {
	return &Handlers{
		state:  state,
		window: window,
	}
}

func (h *Handlers) OnAddVideos() {
	log.Println("Opening file dialog...")
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		log.Println("File dialog callback triggered")
		if err != nil {
			log.Printf("File dialog error: %v", err)
			return
		}
		if reader == nil {
			log.Println("File dialog cancelled")
			return
		}

		log.Println("Getting URI path...")
		uri := reader.URI()
		log.Printf("URI: %v", uri)
		path := uri.Path()
		log.Printf("Path: %s", path)
		
		log.Println("Closing reader...")
		reader.Close()
		log.Println("Reader closed")

		if isVideoFile(path) {
			log.Printf("Adding video: %s", path)
			h.state.AddVideo(path)
			log.Println("Video added")
		}
	}, h.window)

	fd.SetFilter(&videoFilter{})
	log.Println("Showing file dialog...")
	fd.Show()
	log.Println("File dialog shown")
}

func (h *Handlers) OnRemove() {
	h.state.RemoveSelected()
}

func (h *Handlers) OnMoveUp() {
	h.state.MoveUp()
}

func (h *Handlers) OnMoveDown() {
	h.state.MoveDown()
}

func (h *Handlers) OnClear() {
	if h.state.Count() == 0 {
		return
	}

	dialog.ShowConfirm("Clear All", "Remove all videos from the list?", func(ok bool) {
		if ok {
			h.state.Clear()
		}
	}, h.window)
}

type videoFilter struct{}

func (f *videoFilter) Matches(uri fyne.URI) bool {
	return isVideoFile(uri.Path())
}

func (f *videoFilter) Extensions() []string {
	return videoExtensions
}
