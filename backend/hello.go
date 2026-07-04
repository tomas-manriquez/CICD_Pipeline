package main

import (
	"flag"
	"fmt"
)

func main() {
	//declare flag variable
	//path -> String
	var path string
	flag.StringVar(&path, "path", "", "path to golang repository")

	//read flags from command
	flag.Parse()

	//if no flag found, read cargs
	//args -> List
	var args []string
	if path == "" {
		args = flag.Args()
		if len(args) != 1 {
			fmt.Println("incorrect number of arguments passed!")
			fmt.Println("command line arguments: ", flag.Args())
			fmt.Println("ERR -> EXIT")
			return
		} else {
			path = args[0]
		}
	}

	//print to STDOUT
	if path == "" {
		fmt.Println("no path found!")
		fmt.Println("command line arguments: ", flag.Args())
		fmt.Println("ERR -> EXIT")
		return
	} else {
		fmt.Println("path: ", path)
	}
}
