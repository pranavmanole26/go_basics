package main

import (
	"fmt"
	"os"
	"strconv"
)

func getResult(i, j int, o string) int {
	switch o {
	case "*":
		return i * j
	case "/":
		if j == 0 || i == 0 {
			return 0
		}
		return i / j
	case "+":
		return i + j
	case "-":
		return i - j
	case "%":
		if j == 0 || i == 0 {
			return 0
		}
		return i % j
	default:
		fmt.Println("Invalid operation...")
		return 0
	}
}

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Printf("Invalid number of arguments")
		return
	} else {
		o := args[1]                   // operator
		st, e := strconv.Atoi(args[2]) // size of the table

		if e != nil {
			fmt.Println("Invalid argument for size of the table.")
		}
		fmt.Printf("%5s", o)
		for i := 0; i <= st; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Println()
		for j := 0; j <= st; j++ {
			fmt.Printf("%5d", j)
			for i := 0; i <= st; i++ {
				fmt.Printf("%5d", getResult(j, i, o))
			}
			fmt.Println()
		}
	}
}
