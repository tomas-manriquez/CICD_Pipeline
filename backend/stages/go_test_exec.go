package stages

import (
	"errors"
	"fmt"
	"os/exec"
)

func ExecGoTest(path string) bool {
	goTestCmd := exec.Command("go", "test", path)

	goTestOut, err := goTestCmd.Output()

	if err != nil {
		if e, ok := errors.AsType[*exec.Error](err); ok {
			fmt.Println("failed executing: ", e)
		} else if e, ok := errors.AsType[*exec.ExitError](err); ok {
			exitCode := e.ExitCode()
			fmt.Println("command exit rc= ", exitCode)
		} else {
			panic(err)
		}
	}

	fmt.Println(string(goTestOut))

	return true
}
