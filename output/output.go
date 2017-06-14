package output

import (
	"fmt"
	"os"
)

var PrintToStdout = true

func Stdout(args ...interface{}) {
	if PrintToStdout {
		fmt.Fprint(os.Stdout, args...)
	}
}

func StdoutForce(args ...interface{}) {
	fmt.Fprint(os.Stdout, args...)
}

func Fatal(args ...interface{}) {
	fmt.Fprint(os.Stderr, args...)
	os.Exit(1)
}
