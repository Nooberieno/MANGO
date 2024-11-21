package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Gets the users preferred shell for when running commands with the -sh and --shell flags
// Returns the path to the shell binary
func get_shell() string {
	shell := os.Getenv("SHELL")
	fmt.Printf("Using shell: %s to execute commands\n", shell)
	return shell
}

// Executes commands in the shell
// Used when the -sh or --shell flag is set
// Returns an error if a command throws an error, otherwise returns nil
func shell_command(commands []string) error {
	count := 0
	shell := get_shell()
	for _, command := range commands {
		cmd := exec.Command(shell, "-c", command)
		output, err := cmd.CombinedOutput()
		if verbose {
			log.Printf("Executing command: %s\n", command)
		}
		if err != nil {
			return err
		}
		trimmed_output := strings.TrimSpace(string(output))
		if trimmed_output == "" {
			count += 1
			if verbose {
				log.Printf("Command executed succesfully without error")
			}
		} else if !quiet {
			log.Printf("Shell command output: %s\n", trimmed_output)
		}
	}
	if count == len(commands) && !quiet {
		log.Printf("All commands executed succesfully without output\n")
	} else if !quiet {
		num_success := len(commands) - count
		log.Printf("%d commands executed succesfully without output\n", num_success)
	}
	return nil
}

// Executes commands by directly invoking the OS
// When the -sh or --shell flag is set use shell_command instead
// Returns an error if an executed command throws an error, otherwise returns nil
func command(commands []string) error {
	count := 0
	for _, command := range commands {
		cmd_parts := strings.Fields(command)
		if len(cmd_parts) < 1 {
			continue
		}
		cmd := exec.Command(cmd_parts[0], cmd_parts[1:]...)
		output, err := cmd.CombinedOutput()
		if verbose {
			log.Printf("Executing command: %s\n", command)
		}
		if err != nil {
			return err
		}
		trimmed_output := strings.TrimSpace(string(output))
		if trimmed_output == "" {
			count += 1
			if verbose {
				log.Printf("Command succesfully executed without error\n")
			}
		} else if !quiet {
			log.Printf("Command output: %s\n", trimmed_output)
		}
	}
	if count == len(commands) && !quiet {
		log.Printf("All commands executed succesfully without output\n")
	} else if !quiet {
		num_success := len(commands) - count
		log.Printf("%d commands executed succesfully without output\n", num_success)
	}
	return nil
}
