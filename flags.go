package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	verbose    bool
	shell      bool
	quiet      bool
	logfile    bool
	logfileout string
)

func init() {
	flag.BoolVar(&verbose, "v", false, "")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&shell, "sh", false, "")
	flag.BoolVar(&shell, "shell", false, "Enable shell command execution")
	flag.BoolVar(&quiet, "q", false, "")
	flag.BoolVar(&quiet, "quiet", false, "Enable quiet output")
	flag.BoolVar(&logfile, "l", false, "")
	flag.BoolVar(&logfile, "log", false, "Log output to file")
	flag.StringVar(&logfileout, "lf", "MANGO.log", "")
	flag.StringVar(&logfileout, "logfile", "MANGO.log", "Specify output to log MANGO's output, MANGO.log by default")
}

func parse_flags() {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "-shell", "--sh":
				fmt.Fprintf(os.Stderr, "Please use --shel or -s to enable shell command execution, use -h or --help to see all options")
			case "--v", "-verbose":
				fmt.Fprintf(os.Stderr, "Please use --verbose or -v to enable verbose output, use -h or --help to see all options")
			case "-quiet", "--q":
				fmt.Fprintf(os.Stderr, "Please use --quiet or -q to enable quiet output, use -h or --help to see all options")
			case "-help", "--h":
				fmt.Fprintf(os.Stderr, "Please use -h or --help to see all options")
			case "-log", "--l":
				fmt.Fprintf(os.Stderr, "Please use -l or --log to log a the output to a file (MANGO.log) by default")
			case "-logfile", "--lf":
				fmt.Fprintf(os.Stderr, "Please use -lf <filename> or --logfile <filename> to specify a logfile")
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
		fmt.Fprintf(os.Stderr, "-l, --logfile\n\t Log output to a file, MANGO.log file by default\n")
		fmt.Fprintf(os.Stderr, "-lf <filename>, --logfile <filename>, specify a file to log MANGO's output too, MANGO.log when not specified\n")
		fmt.Fprintf(os.Stderr, "-h, --help\n\t List all possible the usage with all possible flags\n")
	}
	flag.Parse()
}
