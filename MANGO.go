package main

import (
	"errors"
	"fmt"
	"os"
)

func CheckBuildFile() {
	if _, err := os.Stat("mango.build"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No mango.build file provided")
		os.Exit(1)
	}
}

func BuildGen(program_name string, lang_name string, comp_name string) {
	file, err := os.Create("mango.build")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	write_string := "Name: " + program_name + "\nLanguage: " + lang_name + "\nCompiler: " + comp_name + "\n"
	_, err = file.WriteString(write_string)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Succesfully written to mango.build file")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}
	if os.Args[1] == "build" {
		CheckBuildFile()
		fmt.Println("mango.build file provided")
	}
	if os.Args[1] == "make" {
		if len(os.Args) < 5 {
			fmt.Println("Missing arguments, mango make <name> <programming language> <compiler>")
		}
		BuildGen(os.Args[2], os.Args[3], os.Args[4])
	}
}
