package main

import (
	"cicd_pipeline/utils"
	"flag"
	"fmt"
)

func main() {
	//declare flag variable
	//path -> String
	var path string
	flag.StringVar(&path, "path", "", "path to golang repository")

	//read flags and argsfrom command
	flag.Parse()
	args := flag.Args()

	//parse args
	path = utils.ParseArgs(path, args)

	//print to STDOUT
	if path == "" {
		fmt.Println("no path found!")
		fmt.Println("command line arguments: ", args)
		fmt.Println("ERR -> EXIT")
		return
	} else {
		fmt.Println("path: ", path)
	}
}

