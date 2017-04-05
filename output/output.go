package output

import (
	"log"
	"os"
)

var PrintToStdout = true
var stdoutLogger = log.New(os.Stdout, "", 0)
var stderrLogger = log.New(os.Stderr, "", 0)

func Stdout(args ...interface{}) {
	if PrintToStdout {
		stdoutLogger.Print(args...)
	}
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
