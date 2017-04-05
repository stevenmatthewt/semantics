package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func runCommand(cmd *exec.Cmd) (string, error) {
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return strings.TrimSpace(out.String()), err
}
