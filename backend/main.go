package main

import (
	"cicd_pipeline/stages"
	"cicd_pipeline/utils"
	"errors"
	"flag"
	"fmt"
)

func main() {
	//declare flag variable
	//path -> String

	//read flags and argsfrom command
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}

func Run() error {
	//read flag, args
	//if no path found, return error
	var path string
	flag.StringVar(&path, "path", "", "path to golang repository")
	flag.Parse()

	args := flag.Args()
	path = utils.ParseArgs(path, args)

	//return error if no path found
	if path == "" {
		fmt.Println("command line arguments: ", args)
		fmt.Println("path: ", path)
		return errors.New("No repository path provided")
	}
	fmt.Println("path: ", path)

	//run stage 1: go test <path>
	stages.ExecGoTest(path)

	return nil
}
