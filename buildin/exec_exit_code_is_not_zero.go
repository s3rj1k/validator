package buildin

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/s3rj1k/validator"
)

// ExecExitCodeIsNotZeroError is a function that defines default error message returned by ExecExitCodeIsNotZero validator.
// nolint: gochecknoglobals
var ExecExitCodeIsNotZeroError = func(v *ExecExitCodeIsNotZero) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' after execution command must exit with none-zero status code", strings.Join(v.Command, " "))
}

// ExecExitCodeIsNotZero is a validator object.
// Validate adds an error if the Command value after execution returns zero exit code.
type ExecExitCodeIsNotZero struct {
	Name           string
	Command        []string
	Message        string
	TimeoutSeconds int64
}

// Validate adds an error if the Command value after execution returns zero exit code.
func (v *ExecExitCodeIsNotZero) Validate(e *validator.Errors) {
	var timeout time.Time

	if v.TimeoutSeconds == 0 {
		timeout = time.Now().Add(
			DefaultExecTimeoutSec * time.Second,
		)
	} else {
		timeout = time.Now().Add(
			time.Duration(v.TimeoutSeconds) * time.Second,
		)
	}

	ctx, cancel := context.WithDeadline(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, v.Command[0], v.Command[1:]...)

	// syscall magic
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pdeathsig: syscall.SIGKILL,
	}

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		e.Add(v.Name, ExecExitCodeIsNotZeroError(v))
		return
	}

	if err == nil {
		e.Add(v.Name, ExecExitCodeIsNotZeroError(v))
	}
}
