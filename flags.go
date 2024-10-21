package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	verbose bool
	shell   bool
	quiet   bool
	logfile bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&shell, "s", false, "")
	flag.BoolVar(&shell, "shell", false, "Enable shell command execution")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&quiet, "quiet", false, "Enable quiet output")
	flag.BoolVar(&logfile, "l", false, "")
	flag.BoolVar(&logfile, "log", false, "Log output to file")
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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: MANGO [flags] <target>\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		fmt.Fprintf(os.Stderr, "-v, --verbose\n\t Enable verbose output\n")
		fmt.Fprintf(os.Stderr, "-s, --shell\n\t Enable shell command execution\n")
		fmt.Fprintf(os.Stderr, "-q, --quiet\n\t Enable quiet output\n")
		fmt.Fprintf(os.Stderr, "-l, --logfile\n\t Log output to MANGO.log file\n")
	}
	flag.Parse()
}
