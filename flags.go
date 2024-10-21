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
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&shell, "s", false, "")
	flag.BoolVar(&shell, "shell", false, "Enable shell command execution")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&quiet, "quiet", false, "Enable quiet output")
}

func parse_flags() {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-shell", "--s":
				log.Fatal("Please use --shel or -s to enable shell command execution, use -h or --help to see all options")
			case "--v", "-verbose":
				log.Fatal("Please use --verbose or -v to enable verbose output, use -h or --help to see all options")
			case "-quiet", "--q":
				log.Fatal("Please use --quiet or -q to enable quiet output, use -h or --help to see all options")
			case "-help", "--h":
				log.Fatal("Please use -h or --help to see all options")
			default:
				continue
			}
		}
		continue
	}
	flag.Parse()
}
