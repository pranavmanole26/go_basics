package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Time Multiplier
//
//  You should get an argument from the command-line and
//  you need to multiply the time duration value `t` with
//  the given argument.
//
//  1. Get an argument from the command-line
//  2. Convert it to int64 and store it in a variable
//  3. Multiply the `t` variable with that variable
//  4. Print the result
//
// HINT
//  You can use ParseInt to convert the command-line
//    argument into an int64 value.
//
//  You can skip the error value using a blank-identifier.
//
// EXPECTED OUTPUT
//
//  When runned like this:
//    go run main.go 2
//
//  It should print this:
//    3h0m0s
// ---------------------------------------------------------

func main() {
	// DONT TOUCH THIS
	// 1h30m means: 1 hour 30 minutes
	t, _ := time.ParseDuration("1h30m")

	val, _ := strconv.ParseInt(os.Args[1], 10, 64)
	t = t * time.Duration(val)

	// DONT TOUCH THIS
	fmt.Println(t)
}
