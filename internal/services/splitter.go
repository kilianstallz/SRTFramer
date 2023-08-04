package services

import (
	"fmt"
	"os"
	"os/exec"
)

type SplitterService struct {
}

func NewSplitterService() SplitterService {

	return SplitterService{}
}

func (s SplitterService) ExtractFrames(filepath string) error {
	fmt.Println("frames directory is empty")
	fmt.Println("extracting frames from video")
	fmt.Println()

	c := exec.Command(
		"ffmpeg", "-i", filepath,
		"-vf", "select='eq(pict_type, I)'", "-vsync", "vfr", "-frame_pts", "true", "frames/%03d.jpg",
	)
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
