//go:build !windows

package server

import (
	"errors"
	"os/exec"
	"syscall"
)

func setNewProcessGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}
}

func sendCtrlBreak(_ int) error {
	return errors.ErrUnsupported
}
