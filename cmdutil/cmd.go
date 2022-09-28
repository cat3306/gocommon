package utils

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"
)

const (
	bash = "/bin/bash"
)

func Cmd(arg string, timeoutArgs ...time.Duration) (string, error) {
	timeout := 3 * time.Second
	if len(arg) == 0 {
		return "", errors.New("arg empty")
	}
	if len(timeoutArgs) > 0 {
		timeout = timeoutArgs[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	args := []string{"-c"}
	args = append(args, arg)
	cmd := exec.CommandContext(ctx, bash, args...)
	raw, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(raw)), err
}
