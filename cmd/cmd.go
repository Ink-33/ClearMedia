// Package cmd is the entry point of the application.
package cmd

import (
	"fmt"
	"os"

	"github.com/Ink-33/ClearMedia/internal/factory"
	"github.com/Ink-33/ClearMedia/utils"
)

const (
	inputDir = iota + 1
	outputDir
)

// Main is the entry point of the application.
func Main() error {
	infiles, err := utils.GetFileList(inputDir)
	if err != nil {
		return err
	}
	err = factory.Do(infiles, os.Args[outputDir], factory.DoAudioStrip)
	if err != nil {
		return err
	}
	fmt.Println("Done")
	return nil
}
