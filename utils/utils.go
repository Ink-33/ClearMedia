// Package utils provides some utility functions.
package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

// FFmpegPath returns the path of the ffprobe executable
var FFmpegPath string

// GetFileList returns a list of files in a directory
func GetFileList(dir int) ([]string, error) {
	files := make([]string, 0)

	path := os.Args[dir]
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	list, err := file.Readdir(0)
	if err != nil {
		return nil, err
	}
	for k := range list {
		if list[k].IsDir() {
			continue
		}
		files = append(files, list[k].Name())
	}
	return files, nil
}

// GetExecutableName returns the name of the executable
func GetExecutableName() string {
	return filepath.Base(os.Args[0])
}

// GetFFmpegPath returns the path of the ffprobe executable
func GetFFmpegPath() (path string) {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		return ""
	}
	return path
}
