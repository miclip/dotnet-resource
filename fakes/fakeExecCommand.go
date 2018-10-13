package fakes

import (
	"fmt"
	"testing"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var MockedExitStatus = 0
var MockedStdout string

// CommandString Command being executed
var CommandString string

func FakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestExecCommandHelper", "--", command}
	cs = append(cs, args...)
	CommandString = command +" " + strings.Join(args, " ")
	cmd := exec.Command(os.Args[0], cs...)
	es := strconv.Itoa(MockedExitStatus)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1",
			"STDOUT=" + MockedStdout,
			"EXIT_STATUS=" + es}
	return cmd
}

func TestExecCommandHelper(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
			return
	}
	fmt.Fprintf(os.Stdout, os.Getenv("STDOUT"))
	i, _ := strconv.Atoi(os.Getenv("EXIT_STATUS"))
	os.Exit(i)
}
