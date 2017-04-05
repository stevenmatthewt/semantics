package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func runCommand(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%v: %s", err, stderr.String())
	}

	return strings.TrimSpace(out.String()), err
}
