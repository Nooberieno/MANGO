package main

import "fmt"

func main() {
	if err := parse(open_build()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("targets:")
		for _, target := range targets {
			fmt.Println(target.Name)
			for _, command := range target.Commands {
				fmt.Println(" -", command)
			}
		}
		fmt.Println("variables:")
		for _, varia := range variables {
			fmt.Println(varia)
		}
	}
}
