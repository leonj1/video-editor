package app

import (
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os/exec"
)

func ExtractThumbnail(videoPath string) (image.Image, error) {
	cmd := exec.Command("ffmpeg",
		"-i", videoPath,
		"-vframes", "1",
		"-f", "image2pipe",
		"-vcodec", "png",
		"-vf", "scale=120:-1",
		"-")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	img, _, err := image.Decode(&out)
	if err != nil {
		return nil, err
	}

	return img, nil
}
