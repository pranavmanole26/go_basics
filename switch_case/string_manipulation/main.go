package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) > 1 {

		c := args[1]
		s := args[2]

		switch c {
		case "lower":
			fmt.Println(strings.ToLower(s))
		case "upper":
			fmt.Println(strings.ToUpper(s))
		case "title":
			fmt.Println(strings.Title(s))
		default:
			fmt.Println("Invalid command...")
		}
	} else {
		fmt.Printf("Too few args. Expected 2 got %d\f", len(args)-1)
	}
}
