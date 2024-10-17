package main

import (
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
	if len(os.Args) >= 2 {
		actual_target := check_target(os.Args[1])
		if actual_target == nil {
			log.Fatal("Unknown target: ", os.Args[1])
		} else {
			err := shell_command(actual_target.Commands)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal("Missing build target, please supply a target to build")
	}
	if err := parse_file(); err != nil {
		log.Fatal(err)
	}
}
