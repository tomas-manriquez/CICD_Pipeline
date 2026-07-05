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

/*
Scenario 01
Given user compiled the backend
When the user executes the program with the --path flag with a valid string
Then the program assigns the value to var path and returns
*/
func TestPathFromFlag(t *testing.T) {
	path:= "~/"
	args:= []string{"a", "b", "c"}
	want:= "~/"
	got := utils.ParseArgs(path, args)
	if got != want {
		t.Errorf(`ParseArgs("~/", []) = %q, want match for %#q`, got, want)
	}
}

/*
Scenario 02
Given user compiled the backend
When the user executes the program with no --path flag but args has 1 valid string
Then the program assigns the value to var path and returns
*/
func TestPathFromArgs(t *testing.T) {
	path:= ""
	args:= []string{"~/"}
	want:= "~/"
	got := utils.ParseArgs(path, args)
	if got != want {
		t.Errorf(`ParseArgs("", ["~/"]) = %q, want match for %#q`, got, want)
	}
}

/*
Scenario 03
Given user compiled the backend
When the user executes the program with no --path and no args
Then the program keeps empty string to var path and returns
*/
func TestNoPath(t *testing.T) {
	path:= ""
	args:= []string{}
	want:= ""
	got:= utils.ParseArgs(path, args)
	if got != want {
		t.Errorf(`ParseArgs("", []) = %q, want match for %#q`, got, want)
	}
}

/*
Scenario 04
Given user compiled the backend
When the user executes the program with no --path but args has 2 or more values
Then the program keeps empty string to var path and returns
*/
func TestTooManyArgs(t *testing.T) {
	path:= ""
	args:= []string{"./", "a"}
	want:= ""
	got:= utils.ParseArgs(path, args)
	if got != want {
		t.Errorf(`ParseArgs("", ["./","a"]) = %q, want match for %#q`, got, want)
	}
}

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
