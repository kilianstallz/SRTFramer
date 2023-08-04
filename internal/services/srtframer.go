package services

import (
	"errors"
	"fmt"
	"github.com/barasher/go-exiftool"
	"path/filepath"
	"regexp"
	"srtframer/internal/models"
	"strconv"
)

type Srtframer struct {
	splitter SplitterService
	files    FileService
}

// TODO: File service should be repository

func NewSrtframer(splitter SplitterService, files FileService) *Srtframer {
	return &Srtframer{
		splitter: splitter,
		files:    files,
	}
}

func (s *Srtframer) Execute(videoPath, strPath string) error {
	err := s.files.CreateOutDir()
	if err != nil {
		return err
	}

	err = s.splitter.ExtractFrames(videoPath)
	if err != nil {
		return err
	}

	files, err := s.files.ReadOutputFiles()
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return errors.New("output directory is empty")
	}

	var frameMap = make(map[int]models.Frame)

	// read all filenames
	for _, file := range files {
		//fmt.Println(file.Name())
		reg := regexp.MustCompile(`([0-9]+).jpg`)
		frame := reg.FindStringSubmatch(file.Name())
		nr, _ := strconv.Atoi(frame[1])
		fr := models.Frame{Number: nr}
		frameMap[nr] = fr
	}

	subtitles, err := s.files.ReadSrtFile(strPath)
	if err != nil {
		return err
	}
	if len(subtitles) == 0 {
		return errors.New("subtitle file is empty")
	}
	fmt.Println("Number of subtitles: ", len(subtitles))

	// TODO: Extract into service
	for _, subtitle := range subtitles {
		nr := subtitle.Number
		reglat := regexp.MustCompile(`\[latitude: ([0-9.]+)\]`)
		reglong := regexp.MustCompile(`\[longitude: ([0-9.]+)\]`)
		regalt := regexp.MustCompile(`\[altitude: ([0-9.]+)\]`)
		lat := reglat.FindStringSubmatch(subtitle.Text)
		long := reglong.FindStringSubmatch(subtitle.Text)
		alt := regalt.FindStringSubmatch(subtitle.Text)

		// check if frame number exists in map
		if _, ok := frameMap[nr-1]; ok {
			//fmt.Printf("%v\n", frameMap[nr])
			frameMap[nr-1] = models.Frame{Number: nr - 1, Lat: lat[1], Long: long[1], Alt: alt[1]}
		}
	}

	fmt.Println("Adding exif metadata")
	// TODO: Into service
	e, _ := exiftool.NewExiftool()

	for _, file := range files {
		path := filepath.Join("frames", file.Name())
		originals := e.ExtractMetadata(path)
		// parse number from filename
		reg := regexp.MustCompile(`([0-9]+).jpg`)
		frameNr := reg.FindStringSubmatch(file.Name())
		nr, _ := strconv.Atoi(frameNr[1])
		frame := frameMap[nr]
		originals[0].SetString("GPSLatitude", frame.Lat)
		originals[0].SetString("GPSLongitude", frame.Long)
		originals[0].SetString("GPSAltitude", frame.Alt)
		originals[0].SetString("GPSLatitudeRef", "N")
		originals[0].SetString("GPSLongitudeRef", "E")
		originals[0].SetString("GPSAltitudeRef", "0")

		// write metadata to image
		e.WriteMetadata(originals)
	}

	err = e.Close()
	if err != nil {
		return err
	}

	return nil

}
