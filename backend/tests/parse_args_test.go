package tests

import (
	"testing"
	"cicd_pipeline/utils"
)

//Feature: Read path to repo from input
//I should write one test that loops every scenario??
/*
tests =

flag provided

flag empty, args contains one

flag empty, args empty

flag empty, args >1

flag provided + args also provided
*/
type testCase struct {
	name string

	path string
	args []string

	want string
}

func TestParseArgs (t *testing.T) {
	tests := []testCase {
		{
			name: "path provided via flag",
			path: "~/",
			args: []string{},
			want: "~/",
		},
		{
			name: "path provided as positional argument",
			path: "",
			args: []string{"~/"},
			want: "~/",
		},
		{
			name: "path not provided neither in flag nor in positional arguments",
			path: "",
			args: []string{},
			want: "",
		},
		{
			name: "path provided as positional arguments with too many arguments",
			path: "",
			args: []string{"./","a"},
			want: "",
		},
		{
			name: "path provided both as flag and positional argument; flag is taken",
			path: "~/",
			args: []string{"./"},
			want: "~/",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := utils.ParseArgs(tc.path, tc.args)

			if got != tc.want {
				t.Errorf(`ParseArgs(%q, %v) = %q, want %q`,
					tc.path,
					tc.args,
					got,
					tc.want,
					)
			}
		})
	}
}

/*
Scenario 01
Given user compiled the backend
When the user executes the program with the --path flag with a valid string
Then the program assigns the value to var path and returns
*/

/*
Scenario 02
Given user compiled the backend
When the user executes the program with no --path flag but args has 1 valid string
Then the program assigns the value to var path and returns
*/

/*
Scenario 03
Given user compiled the backend
When the user executes the program with no --path and no args
Then the program keeps empty string to var path and returns
*/

/*
Scenario 04
Given user compiled the backend
When the user executes the program with no --path but args has 2 or more values
Then the program keeps empty string to var path and returns
*/

/*
Scenario 05
Given user compiled the backend
When the user executes the program with a --path flag value and args has values
Then the program assigns to path the value in the --path flag
*/
func TestFlagBeforeArgs(t *testing.T) {
	path:= "~/"
	args:= []string{"./"}
	want:= "~/"
	got:= utils.ParseArgs(path, args)
	if got != want {
		t.Errorf(`ParseArgs("~/", ["./"]) = %q, want match for %#q`, got, want)
	}
}
