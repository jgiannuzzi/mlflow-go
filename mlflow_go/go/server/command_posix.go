//go:build !windows

package server

import (
	"fmt"
	"os/exec"
	"syscall"
)

func setNewProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}
}

func terminateProcessGroup(cmd *exec.Cmd) error {
	if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to terminate process group: %w", err)
	}

	return nil
}
