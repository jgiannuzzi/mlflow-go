//go:build windows

package server

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"

	"github.com/mlflow/mlflow-go/mlflow_go/go/config"
)

func launchCommand(ctx context.Context, cfg *config.Config) error {
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return fmt.Errorf("could not create job object: %w", err)
	}

	var info windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION
	info.BasicLimitInformation.LimitFlags = windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE
	if _, err := windows.SetInformationJobObject(job, windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info))); err != nil {
		return fmt.Errorf("could not set job object information: %w", err)
	}

	//nolint:gosec
	cmd := exec.CommandContext(ctx, cfg.PythonCommand[0], cfg.PythonCommand[1:]...)
	cmd.Env = append(os.Environ(), cfg.PythonEnv...)
	cmd.Stdout = logrus.StandardLogger().Writer()
	cmd.Stderr = logrus.StandardLogger().Writer()
	cmd.WaitDelay = 5 * time.Second //nolint:mnd
	cmd.Cancel = func() error {
		logrus.Debug("Sending termination signal to command")

		return windows.CloseHandle(job)
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}

	logrus.Debugf("Launching command: %v", cmd)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("command could not launch: %w", err)
	}

	hProc, err := windows.OpenProcess(2097151, true, uint32(cmd.Process.Pid))
	if err != nil {
		return fmt.Errorf("could not open process: %w", err)
	}
	defer windows.CloseHandle(hProc)

	if err := windows.AssignProcessToJobObject(job, hProc); err != nil {
		return fmt.Errorf("could not assign process to job object: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command exited with error: %w", err)
	}

	return nil
}
