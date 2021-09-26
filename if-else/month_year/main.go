package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	start, end time.Time
)

func isLeap(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}

func timeTaken() {
	end = time.Now()
	fmt.Printf("End time: %v.\n", end)
	fmt.Printf("Time taken: %v.\n", end.Sub(start))
}

func main() {
	start = time.Now()
	fmt.Printf("Start time: %v.\n", start)
	defer timeTaken()

	var (
		args       []string
		month      string
		year, days int
		err        error
	)

	if args = os.Args; len(args) != 3 {
		fmt.Printf("Invalid number of arguments. Expected 2 got %d.\n", len(args)-1)
		return
	}

	month = args[1]
	if year, err = strconv.Atoi(args[2]); err != nil {
		fmt.Printf("Invalid year: %q", year)
	} else {
		switch month {
		case "jan", "mar", "may", "jul", "aug", "oct", "dec":
			days = 31
		case "apr", "jun", "sep", "nov":
			days = 30
		case "feb":
			if isLeap(year) {
				days = 28
			} else {
				days = 29
			}
		default:
			fmt.Println("Invalid month")
		}

	}

	fmt.Printf("In month %s and year %d, days are %d\n", month, year, days)
}
