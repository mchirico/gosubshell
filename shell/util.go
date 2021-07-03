package shell

import (
	"github.com/mingrammer/cfmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"strings"
)

func CreateSubShell(cmd *cobra.Command, args []string) {

	// setup sub-shell env vars
	os.Setenv("AWS", "AWS")
	cfmt.Successln("Starting sub-shell. run echo $AWS ")

	// drop into pre-configured shell
	shell := "/usr/bin/env bash"

	shellCmd := strings.Fields(shell)
	subShell := exec.Command(shellCmd[0], shellCmd[1:]...)

	subShell.Stderr = os.Stderr
	subShell.Stdout = os.Stdout
	subShell.Stdin = os.Stdin
	subShell.Run()
}
