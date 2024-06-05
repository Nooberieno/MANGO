package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func command(command string) {
	CmdArgs := strings.Split(command, " ")
	cmd := exec.Command(CmdArgs[0], CmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error runing command: ", err)
		os.Exit(1)
	}
}
