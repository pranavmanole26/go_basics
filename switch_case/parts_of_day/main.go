package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	switch {
	case now.Hour() >= 18:
		fmt.Println("Good Evening...")
	case now.Hour() >= 12:
		fmt.Println("Good Afternoon...")
	case now.Hour() >= 6:
		fmt.Println("Good Morning...")
	default:
		fmt.Println("Good Night...")
	}
}
