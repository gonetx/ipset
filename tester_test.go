package ipset

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var (
	needError bool
	flag      = struct{}{}
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	if needError {
		cmd.Env = append(cmd.Env, "GO_WANT_HELPER_NEED_ERR=1")
	}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}

	if len(args) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "No command")
		os.Exit(2)
	}

	if os.Getenv("GO_WANT_HELPER_NEED_ERR") == "1" {
		_, _ = fmt.Fprintf(os.Stderr, "fake error")
		os.Exit(1)
	}

	if len(args) > 1 {
		switch args[1] {
		case _version:
			if args[0] == "non-supported" {
				_, _ = fmt.Fprintf(os.Stdout, invalidVersion)
			} else {
				_, _ = fmt.Fprintf(os.Stdout, validVersion)
			}
		case _list:
			if findOption(args, "-resolve") {
				_, _ = fmt.Fprintf(os.Stdout, listInfoResolved)
			} else {
				_, _ = fmt.Fprintf(os.Stdout, listInfo)
			}
		case _save:
			if findOption(args, "-resolve") {
				_, _ = fmt.Fprintf(os.Stdout, saveInfoResolved)
			} else {
				_, _ = fmt.Fprintf(os.Stdout, saveInfo)
			}
		case _test:
			if len(args) > 3 && args[3] == testNotExistIp {
				_, _ = fmt.Fprintf(os.Stderr, "1.1.1.2 is NOT in set foo.")
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}

func findOption(args []string, target string) bool {
	for _, arg := range args {
		if arg == target {
			return true
		}
	}
	return false
}

func setupCmd(flag ...struct{}) {
	execCommand = fakeExecCommand
	if len(flag) > 0 {
		needError = true
	}
}

func teardownCmd() {
	execCommand = exec.Command
	needError = false
}

func setupLookPath(filename ...string) {
	fn := ""
	if len(filename) > 0 {
		fn = filename[0]
	}
	execLookPath = func(f string) (s string, err error) {
		switch fn {
		case "error":
			err = fmt.Errorf("path not exist")
		case "":
			s = f
		default:
			s = fn
		}
		return
	}
}

func teardownLookPath() {
	execLookPath = exec.LookPath
	ipsetPath = ""
}

const (
	validVersion   = "ipset v6.29, protocol version: 6"
	invalidVersion = "ipset v5.1, protocol version: 5"
	listInfo       = `
Name: foo
Type: hash:ip
Revision: 4
Header: family inet hashsize 1024 maxelem 65536
Size in memory: 168
References: 0
Number of entries: 1
Members:
1.1.1.1`
	listInfoResolved = `
Name: foo
Type: hash:ip
Revision: 4
Header: family inet hashsize 1024 maxelem 65536
Size in memory: 168
References: 0
Number of entries: 1
Members:
one.one.one.one`
	saveInfo = `
create foo hash:ip family inet hashsize 1024 maxelem 65536
add foo 1.1.1.1
`
	saveInfoResolved = `
create foo hash:ip family inet hashsize 1024 maxelem 65536
add foo one.one.one.one
`
	testNotExistIp = "1.1.1.2"
)
