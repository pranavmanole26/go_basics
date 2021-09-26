package main

import (
	"fmt"
)

func main() {
	var y int
	fmt.Println("Enter year: ")
	fmt.Scanf("%d", &y)

	if y%100 == 0 && y%400 == 0 || y%100 != 0 && y%4 == 0 {
		fmt.Println("Leap")
	} else {
		fmt.Println("Non Leap")
	}

}
