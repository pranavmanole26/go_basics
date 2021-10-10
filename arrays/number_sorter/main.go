package main

import (
	"fmt"
	"os"
	"strconv"
)

func sortAsc(n1, n2 int) (int, int) {
	// fmt.Printf("n1=%d, n2=%d\n", n1, n2)
	if n1 > n2 {
		return n2, n1
	}
	return n1, n2
}

func strToInt64(s string) (n int64) {
	n, e := strconv.ParseInt(s, 10, 32)
	if e != nil {
		fmt.Printf("Parsing error, while parsing %s\nError:%v", s, e)
		n = 0
	}
	return
}

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Too few arguments")
		return
	} else {
		nums := os.Args[1:]
		sortedNums := make([]int, len(nums))
		for i := range nums {
			sortedNums[i] = int(strToInt64(nums[i]))
		}
		// var sortedNums []int = make([]int, len(nums))
		for i := range sortedNums {
			if i > 0 {
				// fmt.Println("new")
				for j := range sortedNums[:i+1] {
					fmt.Printf("%v\n", sortedNums[:i+1])
					sortedNums[j], sortedNums[i] = sortAsc(sortedNums[i], sortedNums[j])
				}
			}
		}
		fmt.Printf("Sorted nums %v", sortedNums)
	}
}
