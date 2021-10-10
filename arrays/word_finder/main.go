package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	corpus      = "lazy cat jumps again and again and again since the beginning this was very important"
	noSearchStr = "and or was the since very"
)

func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Println("Too few arguments. Expected more than 1.")
		return
	} else {
		// strs := strings.Fields(corpus)
		filteredStrs := make([]string, 0)
		for _, s1 := range strings.Fields(corpus) {
			if !strings.Contains(noSearchStr, s1) {
				filteredStrs = append(filteredStrs, s1)
			}
		}
		args := os.Args[1:]
		for i := range args {
			if strings.Contains(strings.ToLower(strings.Join(filteredStrs, " ")), strings.ToLower(args[i])) {
				index := strings.Index(corpus, args[i])
				// fmt.Println(index)
				if index != -1 {
					fmt.Printf("#%3d: %q\n", index, args[i])
				}
			}

		}
	}
}
