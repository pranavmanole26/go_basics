package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3}
	nums := make([]int, 3, 5)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	Show("nums", nums)

	num1 := nums[1:3]
	nums[2] = 6
	Show("num1:", num1)

	nums = append(nums, 4)
	Show("nums", nums)

	nums = append(nums, 7)
	nums = append(nums, 8)
	Show("nums beyond cap:", nums)
}
