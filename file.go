package main

import (
	"log"
	"os"
)

// Checks if the default build file exists if it is continue, if not throw an error
// Returns a string containing the build file path
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
	}
	log.Fatal("No valid build file")
	return ""
}
