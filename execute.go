package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func get_shell() string {
	shell := os.Getenv("SHELL")
	fmt.Println(shell)
	return shell
}

func shell_command(commands []string) error {
	shell := get_shell()
	for _, command := range commands {
		fmt.Println(command)
		cmd := exec.Command(shell, "-c", command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("%v\n", err)
		}
		trimmed_output := strings.TrimSpace(string(output))
		if trimmed_output == "" {
			log.Printf("Command executed succesfully with no output\n")
		} else {
			log.Printf("Command output: %s\n", trimmed_output)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
