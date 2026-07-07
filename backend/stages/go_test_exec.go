package stages

import (
	"errors"
	"fmt"
	"os/exec"
)

/*
The JSON stream is a newline-separated sequence of TestEvent objects
corresponding to the Go struct:

    type TestEvent struct {
        Time        time.Time // encodes as an RFC3339-format string
        Action      string
        Package     string
        Test        string
        Elapsed     float64 // seconds
        Output      string
        FailedBuild string
    }

The Time field holds the time the event happened. It is conventionally omitted
for cached test results.

The Action field is one of a fixed set of action descriptions:

    start  - the test binary is about to be executed
    run    - the test has started running
    pause  - the test has been paused
    cont   - the test has continued running
    pass   - the test passed
    bench  - the benchmark printed log output but did not fail
    fail   - the test or benchmark failed
    output - the test printed output
    skip   - the test was skipped or the package contained no tests

Every JSON stream begins with a "start" event.

The Package field, if present, specifies the package being tested. When the
go command runs parallel tests in -json mode, events from different tests are
interlaced; the Package field allows readers to separate them.

The Test field, if present, specifies the test, example, or benchmark function
that caused the event. Events for the overall package test do not set Test.

The Elapsed field is set for "pass" and "fail" events. It gives the time elapsed
for the specific test or the overall package test that passed or failed.


The Output field is set for Action == "output" and is a portion of the
test's output (standard output and standard error merged together). The
output is unmodified except that invalid UTF-8 output from a test is coerced
into valid UTF-8 by use of replacement characters. With that one exception,
the concatenation of the Output fields of all output events is the exact output
of the test execution.

The FailedBuild field is set for Action == "fail" if the test failure was caused
by a build failure. It contains the package ID of the package that failed to
build. This matches the ImportPath field of the "go list" output, as well as the
BuildEvent.ImportPath field as emitted by "go build -json".
*/

func ExecGoTest(path string) bool {
	goTestCmd := exec.Command("go", "test", path, "-json")

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
