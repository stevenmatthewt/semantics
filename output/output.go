package output

import (
	"log"
	"os"
)

var PrintToStdout = true
var stdoutLogger = log.New(os.Stdout, "", 0)
var stderrLogger = log.New(os.Stderr, "", 0)

// Resolver defines an interface for marking a previous output as
// either failed or succeeded.
//
// At the moment, it only operates on the last output.
type Resolver interface {
	Success()
	Failure()
}

type resolver struct{}

func (r resolver) Success() {
	stdoutLogger.Print(" Success!")
}

func (r resolver) Failure() {
	stdoutLogger.Print(" Failure!")
}

func Stdout(args ...interface{}) resolver {
	if PrintToStdout {
		stdoutLogger.Print(args...)
	}

	return resolver{}
}

func StdoutForce(args ...interface{}) {
	stdoutLogger.Print(args...)
}

func Stderr(args ...interface{}) {
	stderrLogger.Print(args...)
}

func Fatal(args ...interface{}) {
	stderrLogger.Fatal(args...)
}
