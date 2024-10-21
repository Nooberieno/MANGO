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
	quiet   bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "Enable verbose output")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&shell, "s", false, "Enable shell command execution")
	flag.BoolVar(&shell, "shell", false, "Enable shell command execution")
	flag.BoolVar(&quiet, "q", false, "Enable quiet output")
	flag.BoolVar(&quiet, "quiet", false, "Enable quiet output")
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
				log.Fatal("Please use --verbose or -v to enable verbose output")
			case ("--q"):
				log.Fatal("Please use --quiet or -q to enable quiet output")
			case ("-quiet"):
				log.Fatal("Please use --quiet or -q to enable quiet output")
			default:
				continue
			}
		}
		continue
	}
	flag.Parse()
}
