package validators

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/s3rj1k/validator"
)

// ExecExitCodeHasAnyError is a function that defines default error message returned by ExecExitCodeHasAny validator.
// nolint: gochecknoglobals
var ExecExitCodeHasAnyError = func(v *ExecExitCodeHasAny) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' after execution command must exit with any of defined status codes", strings.Join(v.Command, " "))
}

// ExecExitCodeHasAny is a validator object.
// Validate adds an error if the Command value after execution returns with non of defined status codes.
type ExecExitCodeHasAny struct {
	Name           string
	Command        []string
	Message        string
	ExitCodes      []int
	TimeoutSeconds int64
}

// Validate adds an error if the Command value after execution returns with non of defined status codes.
func (v *ExecExitCodeHasAny) Validate(e *validator.Errors) {

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

	_ = cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		e.Add(v.Name, ExecExitCodeHasAnyError(v))
		return
	}

	for _, ec := range v.ExitCodes {
		if ec == cmd.ProcessState.ExitCode() {
			return
		}
	}

	e.Add(v.Name, ExecExitCodeHasAnyError(v))
}
