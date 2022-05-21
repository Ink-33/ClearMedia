package main

import (
	"fmt"
	"os"

	"github.com/Ink-33/ClearMedia/cmd"
	"github.com/Ink-33/ClearMedia/internal/base"
	"github.com/Ink-33/ClearMedia/utils"
)

func main() {
	fmt.Println("ClearMedia Version:", base.Version)
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <input dir> <output dir>\n", utils.GetExecutableName())
		return
	}

	ff := utils.GetFFmpegPath()
	if ff == "" {
		fmt.Printf("Error: FFmpeg not found")
		return
	}
	fmt.Printf("Found FFmpeg at : %s\n", ff)
	utils.FFmpegPath = ff

	err := cmd.Main()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
