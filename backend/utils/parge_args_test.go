package tests

import (
	"testing"
	"cicd_pipeline/utils"
)

//Feature: Read path to repo from input

/*
Scenario 01
Given user compiled the backend
When the user executes the program with the --path flag with a valid string
Then the program assigns the value to var path and returns
*/
func PathFromFlag(t *testing.T) {
	path:= "~/"
	args:= ["a", "b", "c"]
	want:= "Users/tomas"
	path := ParseArgs(path, args)
	if !want.MatchString(path) {
		t.Errorf(`ParseArgs("~/") = %q, want match for %#q`, path, want)
	}
}

/*
Scenario 02
Given user compiled the backend
When the user executes the program with no --path flag but cargs has 1 valid string
Then the program assigns the value to var path and returns
*/
