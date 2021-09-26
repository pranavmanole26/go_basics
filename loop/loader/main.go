package main

import (
	"fmt"
	"time"
)

func main() {
	loader := []string{"\\", "|", "/", "-"}

	for i := 0; ; i++ {
		if i == len(loader) {
			i = 0
		}
		fmt.Printf("\r%s Please Wait. Processing...", loader[i])
		time.Sleep(time.Duration(250 * 1000 * 1000))
	}
}
