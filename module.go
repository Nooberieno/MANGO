package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_modules(dev bool) []string {
	var prefix string
	if dev {
		prefix = "module?dev"
	} else {
		prefix = "module"
	}
	var modules []string
	file, err := os.Open("mango.build")
	if err != nil {
		fmt.Println("Error opening build file: ", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if strings.HasPrefix(line, prefix) {
			ModName := strings.Fields(line)[1]
			ModName = strings.TrimSuffix(ModName, ":")
			modules = append(modules, ModName)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading build file: ", err)
	}
	return modules
}

func contains_module(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
