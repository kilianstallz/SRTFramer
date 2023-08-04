package services

import (
	"errors"
	gosrt "github.com/konifar/go-srt"
	"os"
	"path/filepath"
)

type FileService struct {
	outDirPath string
}

func NewFileService(outPath string) FileService {
	return FileService{
		outDirPath: CreateOutputPath(outPath),
	}
}

func CreateOutputPath(outPath string) string {
	return filepath.Join(outPath, "frames")
}

func (s FileService) CreateOutDir() error {
	_, err := os.Stat(s.outDirPath)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("frames", 0755)
		if errDir != nil {
			panic(err)
		}
	}

	files, err := os.ReadDir(s.outDirPath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return nil
	}

	return errors.New("frames directory not empty")
}

func (s FileService) ReadOutputFiles() ([]os.DirEntry, error) {
	return os.ReadDir(s.outDirPath)
}

func (s FileService) ReadSrtFile(path string) ([]gosrt.Subtitle, error) {
	subtitles, err := gosrt.ReadFile(path)
	if err != nil {
		return make([]gosrt.Subtitle, 0), err
	}

	return subtitles, nil
}
