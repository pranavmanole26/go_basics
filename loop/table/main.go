package main

import (
	"fmt"
)

func main() {
	fmt.Print("    x")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for j := 0; j <= 10; j++ {
		fmt.Printf("%5d", j)
		for i := 0; i <= 5; i++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}

}
