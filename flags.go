package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

var (
	verbose bool
	shell   bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Enable verbose mode")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose mode")
	flag.BoolVar(&shell, "s", false, "Enable shell command execution")
	flag.BoolVar(&shell, "shell", false, "Enable shell command execution")
}

func parse_flags() {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case ("-verbose"):
				log.Fatal("Please use --verbose or -v to enable verbose output")
			case ("-shell"):
				log.Fatal("Please use --shel or -s to enable shell command execution")
			case ("--s"):
				log.Fatal("Please use --shel or -s to enable shell command execution")
			case ("--v"):
				log.Fatal("please use --verbose or -v to enable verbose output")
			default:
				continue
			}
		}
		continue
	}
	flag.Parse()
}
