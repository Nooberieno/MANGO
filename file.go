package main

import (
	"log"
	"os"
)

func check_file() (build_file string) {
	_, ERR := os.Stat("MANGO.build")
	_, err := os.Stat("mango.build")
	if ERR == nil && err == nil && !quiet {
		log.Println("Both build files exist, continuing with MANGO.build")
		return "MANGO.build"
	} else if ERR == nil {
		return "MANGO.build"
	} else if err == nil {
		return "mango.build"
	} else {
		log.Fatal("No valid build file")
	}
	return ""
}
