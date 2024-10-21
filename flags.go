package main

import "flag"

var (
	verbose bool
	shell   bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&shell, "shell", false, "Enable shell command execution")
}

func parse_flags() {
	flag.Parse()
}
