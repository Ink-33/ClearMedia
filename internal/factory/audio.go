package factory

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/Ink-33/ClearMedia/internal/order"
	"github.com/Ink-33/ClearMedia/utils"
	"github.com/pkg/errors"
)

func audioStriper(order order.Order) error {
	args := []string{
		utils.FFmpegPath,
		"-i",
		order.GetInput(),
		"-map",
		"0:0",
		"-map_metadata",
		"-1",
		"-c:a",
		"copy",
		filepath.Join(order.GetOutputDir(), filepath.Base(order.GetInput())),
	}

	errwriter := bytes.NewBuffer(nil)
	cmd := &exec.Cmd{
		Path:   utils.FFmpegPath,
		Args:   args,
		Stdout: nil,
		Stderr: errwriter,
	}
	if cmd.Run() != nil {
		msgs := bytes.Split(errwriter.Bytes(), []byte{'\n'})
		return errors.Errorf("%s", string(msgs[len(msgs)-2]))
	}
	return nil
}
