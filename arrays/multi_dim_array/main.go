package main

import (
	"fmt"
	"strings"
)

func main() {

	scientists := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i, scientist := range scientists {
		fmt.Printf("%-10s %-10s %-10s", scientist[0], scientist[1], scientist[2])
		fmt.Println()
		if i == 0 {
			fmt.Printf("%s", strings.Repeat("=", 35))
			fmt.Println()
		}
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
}
