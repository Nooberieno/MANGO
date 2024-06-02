package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}
	for i, arg := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", i+1, arg)
	}
}
