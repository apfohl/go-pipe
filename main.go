package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Printf("%v, commit %v, built at %v\n", version, commit, date)
		os.Exit(0)
	}

	fmt.Println(fmt.Sprintf("%s\n", Hello("world")))
}

func Hello(value string) string {
	return fmt.Sprintf("Hello %s!", value)
}
