package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check_file() (build_file string) {
	_, ERR := os.Stat("MANGO.build")
	_, err := os.Stat("mango.build")
	if ERR == nil && err == nil {
		fmt.Println("Both build files exist, continuing with MANGO.build")
		return "MANGO.build"
	} else if ERR == nil {
		return "MANGO.build"
	} else if err == nil {
		return "mango.build"
	} else {
		panic("No valid build file")
	}

}

func open_build() (contents string) {
	file, err := os.Open(check_file())
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		} else {
			contents += line + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return contents
}
