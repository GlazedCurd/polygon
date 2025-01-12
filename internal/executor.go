package internal

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

const defautlTimeoutSec = 10

func runCmd(cmd *CommandCfg, in io.Reader, out io.Writer, er io.Writer, logger *log.Logger) error {
	if cmd.Cmd == "" {
		logger.Info("Empty command provided")
		return nil
	}

	logger.Debug("Starting Command")

	timeout := defautlTimeoutSec * time.Second

	if cmd.Timeout.Microseconds() != 0 {
		timeout = cmd.Timeout
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	buildcmd := exec.CommandContext(
		ctx,
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
