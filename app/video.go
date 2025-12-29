package app

import (
	"fmt"
	"os"
	"path/filepath"
)

type Video struct {
	Path string
	Name string
	Size int64
}

func NewVideo(path string) (*Video, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &Video{
		Path: path,
		Name: filepath.Base(path),
		Size: info.Size(),
	}, nil
}

func (v *Video) SizeString() string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case v.Size >= GB:
		return formatSize(float64(v.Size)/GB, "GB")
	case v.Size >= MB:
		return formatSize(float64(v.Size)/MB, "MB")
	case v.Size >= KB:
		return formatSize(float64(v.Size)/KB, "KB")
	default:
		return formatSize(float64(v.Size), "B")
	}
}

func formatSize(size float64, unit string) string {
	if size == float64(int(size)) {
		return fmt.Sprintf("%d %s", int(size), unit)
	}
	return fmt.Sprintf("%.1f %s", size, unit)
}
