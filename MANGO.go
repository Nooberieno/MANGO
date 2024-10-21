package main

import (
	"flag"
	"log"
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
