package internal

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/charmbracelet/log"
)

func runCmd(cmd *CommandCfg, in io.Reader, out io.Writer, er io.Writer, logger *log.Logger) error {
	if cmd.Cmd == "" {
		logger.Info("Empty command provided")
		return nil
	}

	logger.Debug("Starting Command")

	buildcmd := exec.Command(
		cmd.Cmd,
		strings.Split(cmd.Args, " ")...,
	)
	buildcmd.Stdin = in
	buildcmd.Stdout = out
	buildcmd.Stderr = er
	logger.Debugf("Executing %s", buildcmd.String())

	err := buildcmd.Run()
	if err != nil {
		return fmt.Errorf("running command: %s", err)
	}
	return nil
}
