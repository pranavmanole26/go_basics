package main

import "fmt"

func main() {

	var nums [2]int

	nums[0] = 10
	nums[0] = -10
	nums[1] -= 10

	fmt.Printf("%v\n", nums)

}
