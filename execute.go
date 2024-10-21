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
	fmt.Printf("Using shell %s to execute commands\n", shell)
	return shell
}

func shell_command(commands []string) error {
	shell := get_shell()
	for _, command := range commands {
		cmd := exec.Command(shell, "-c", command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		trimmed_output := strings.TrimSpace(string(output))
		if trimmed_output == "" {
			log.Printf("Shell command executed succesfully with no output\n")
		} else {
			log.Printf("Shell command output: %s\n", trimmed_output)
		}
	}
	return nil
}

func command(commands []string) error {
	for _, command := range commands {
		cmd_parts := strings.Fields(command)
		if len(cmd_parts) < 1 {
			continue
		}
		cmd := exec.Command(cmd_parts[0], cmd_parts[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		trimmed_output := strings.TrimSpace(string(output))
		if trimmed_output == "" {
			log.Printf("Command executed succesfully with no output\n")
		} else {
			log.Printf("Command output: %s\n", trimmed_output)
		}
	}
	return nil
}
