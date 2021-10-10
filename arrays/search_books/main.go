package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Checks if main string has substring without checking case
func sameCasein(ms, ss string) bool {
	ms = strings.ToLower(ms)
	ss = strings.ToLower(ss)
	return strings.Contains(ms, ss)
}

func main() {

	args := os.Args

	if len(args) != 2 {
		log.Println("Invalid number of arguments. Required 1.")
	} else {

		books := [4]string{
			"Kafka's Revenge",
			"Stay Golden",
			"Everythingship",
			"Kafka's Revenge 2nd Edition",
		}

		sb := args[1] // search book

		for _, b := range books {
			if sameCasein(b, sb) {
				fmt.Printf("%q\n", b)
			}
		}

	}

}
