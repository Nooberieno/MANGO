package main

import "fmt"

func main() {
	parse(open_build())
	for _, target := range targets {
		fmt.Println(target)
	}
	for _, varia := range variables {
		fmt.Println(varia)
	}
}
