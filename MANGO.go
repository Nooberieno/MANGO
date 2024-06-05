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
	write_string := "all: \n\tName: " + program_name + "\n\tLanguage: " + lang_name + "\n\tCompiler: " + comp_name + "\n\nmodule install:\n\t\nmodule?dev clean: #require #module install\n\t"
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
	switch os.Args[1] {
	case "build":
		CheckBuildFile()
		fmt.Println("mango.build file provided")
	case "init":
		if len(os.Args) < 5 {
			fmt.Println("Missing arguments, MANGO make <name> <programming language> <compiler/interpreter>")
		}
		BuildGen(os.Args[2], os.Args[3], os.Args[4])
	case "check":
		modules := get_modules(false)
		devmodules := get_modules(true)
		for _, v := range modules {
			fmt.Println("Module: ", v)
		}
		for _, v := range devmodules {
			fmt.Println("Developer Module: ", v)
		}
	case "command":
		command("echo $HOME")
	default:
		modules := get_modules(false)
		if contains_module(modules, os.Args[1]) {
			fmt.Println("Executing module: ", os.Args[1])
		} else {
			fmt.Println("Unknown command: ", os.Args[1])
		}
	}
}
