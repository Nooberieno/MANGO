package main

import (
	"flag"
	"log"
	"os"
)

func check_target(target string) *Target {
	for _, item := range targets {
		if item.Name == target {
			return &item
		}
	}
	return nil
}

func main() {
	parse_flags()
	if logfile {
		file, err := os.OpenFile(logfileout, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
	}
	args := flag.Args()
	if err := parse_file(); err != nil {
		log.Fatal(err)
	}
	if len(args) >= 1 {
		actual_target := check_target(args[0])
		if actual_target == nil {
			log.Fatal("Unknown target: ", args[0])
		} else if shell {
			err := shell_command(actual_target.Commands)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := command(actual_target.Commands)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal("Missing build target, please supply a target to build")
	}
}
