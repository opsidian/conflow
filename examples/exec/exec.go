package main

import (
	"bufio"
	"bytes"
	"os/exec"
	"syscall"

	"github.com/opsidian/basil/basil"
	"golang.org/x/xerrors"
)

//go:generate basil generate
type Exec struct {
	id       basil.ID `basil:"id"`
	cmd      string   `basil:"required"`
	params   []string
	dir      string
	env      []string
	exitCode int    `basil:"out"`
	stdout   string `basil:"out"`
	stderr   string `basil:"out"`
}

func (e *Exec) Main(ctx basil.BlockContext) error {
	cmd := exec.CommandContext(ctx.Context(), e.cmd, e.params...)

	if e.dir != "" {
		cmd.Dir = e.dir
	}
	if len(e.env) > 0 {
		cmd.Env = e.env
	}

	stdoutReader, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		return xerrors.Errorf("failed to create stdout reader: %v", stdoutErr)
	}

	stderrReader, stdinErr := cmd.StderrPipe()
	if stdinErr != nil {
		return xerrors.Errorf("failed to create stderr reader: %v", stdinErr)
	}

	if startErr := cmd.Start(); startErr != nil {
		return xerrors.Errorf("Failed to start command: %v", startErr)
	}

	var stdout bytes.Buffer

	go func() {
		scanner := bufio.NewScanner(stdoutReader)
		for scanner.Scan() {
			str := scanner.Text()
			stdout.WriteString(str)
			stdout.WriteString("\n")
		}
	}()

	var stderr bytes.Buffer
	go func() {
		scanner := bufio.NewScanner(stderrReader)
		for scanner.Scan() {
			str := scanner.Text()
			stderr.WriteString(str)
			stderr.WriteString("\n")
		}
	}()

	resultErr := cmd.Wait()

	e.exitCode = 0
	if resultErr != nil {
		e.exitCode = 256 // If we can't get the exit code at least return with a custom one
		if exitErr, ok := resultErr.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				e.exitCode = status.ExitStatus()
			}
		}
	}

	e.stdout = stdout.String()
	e.stderr = stderr.String()

	return nil
}
