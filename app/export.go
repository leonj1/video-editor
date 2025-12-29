package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ExportProgress struct {
	Status string
	Done   bool
	Error  error
}

func ExportVideos(videos []*Video, outputPath string, progress chan<- ExportProgress) {
	defer close(progress)

	if len(videos) == 0 {
		progress <- ExportProgress{Error: fmt.Errorf("no videos to export")}
		return
	}

	progress <- ExportProgress{Status: "Preparing export..."}

	tmpFile, err := os.CreateTemp("", "video-list-*.txt")
	if err != nil {
		progress <- ExportProgress{Error: fmt.Errorf("failed to create temp file: %w", err)}
		return
	}
	defer os.Remove(tmpFile.Name())

	for _, video := range videos {
		escaped := strings.ReplaceAll(video.Path, "'", "'\\''")
		fmt.Fprintf(tmpFile, "file '%s'\n", escaped)
	}
	tmpFile.Close()

	progress <- ExportProgress{Status: "Combining videos..."}

	ext := strings.ToLower(filepath.Ext(outputPath))
	args := []string{
		"-f", "concat",
		"-safe", "0",
		"-i", tmpFile.Name(),
	}

	if ext == ".mp4" || ext == ".mov" || ext == ".m4v" {
		args = append(args, "-c", "copy", "-movflags", "+faststart")
	} else {
		args = append(args, "-c", "copy")
	}

	args = append(args, "-y", outputPath)

	cmd := exec.Command("ffmpeg", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		progress <- ExportProgress{Error: fmt.Errorf("ffmpeg error: %w\n%s", err, string(output))}
		return
	}

	progress <- ExportProgress{Status: "Export complete!", Done: true}
}
