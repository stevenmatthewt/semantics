package output

import (
	"fmt"
	"os"
)

var PrintToStdout = true

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
	fmt.Fprint(os.Stdout, " Success!\n")
}

func (r resolver) Failure() {
	fmt.Fprint(os.Stdout, " Failure!\n")
}

func Stdout(args ...interface{}) resolver {
	if PrintToStdout {
		fmt.Fprint(os.Stdout, args...)
	}

	return resolver{}
}

func StdoutForce(args ...interface{}) {
	fmt.Fprint(os.Stdout, args...)
}

func Stderr(args ...interface{}) {
	fmt.Fprint(os.Stderr, args...)
}

func Fatal(args ...interface{}) {
	fmt.Fprint(os.Stderr, args...)
	os.Exit(1)
}
