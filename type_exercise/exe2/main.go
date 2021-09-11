package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// 1. EXERCISE: Refactor Feet to Meter
//
//  Define your own Feet and Meters types.
//
//  Follow the steps inside the code.
// ---------------------------------------------------------

// func main() {
// 	// ----------------------------
// 	// 1. Define Feet and Meters types below
// 	//    Their underlying type can be float64
// 	// ...
// 	type Feet float64
// 	type Meters float64
// 	// ----------------------------
// 	// 2. UNCOMMENT THE CODE BELOW THEN DON'T TOUCH IT
// 	var (
// 		feet   Feet
// 		meters Meters
// 	)

// 	// ----------------------------
// 	// 3. Get feet value from the command-line
// 	// 4. Convert it to an float64 first using ParseFloat
// 	// 5. Then, convert it into a Feet type
// 	fmt.Println("Enter length in feet")
// 	fmt.Scanf("%f", &feet)

// 	// ----------------------------
// 	// 6. Uncomment the code below
// 	// 7. And, convert the expression to Meters type

// 	meters = Meters(float64(feet) * 0.3048)

// 	// ----------------------------
// 	// 8. UNCOMMENT THE CODE BELOW
// 	fmt.Printf("%g feet is %g meters.\n", feet, meters)
// }

// 2. EXERCISE: Parse Arg Numbers
//
//  Use strconv.ParseInt function to get int8, int16, and
//  int32, and int64 values from command-line.
//
// HINT
//  The third argument to ParseInt function represents
//  the bitsize.
//
//  So, giving it 8 returns an int8 convertable value;
//  whereas 16 returns an int16 convertable value.
//
//  Please explore the documentation of ParseInt function
//  and learn how it works.
//
// EXPECTED OUTPUT
//   When runned like this:
//     go run main.go 50 25000 2000000 50000000000000000 00000100
//
//   It should return this:
//     int8 value is : 50
//     int16 value is: 25000
//     int32 value is: 2000000
//     int64 value is: 50000000000000000
//     00000100 is: 4
// ---------------------------------------------------------

func main() {
	// --------------------------------------
	// EXAMPLE:
	// --------------------------------------
	// How to get an int8 from command-line:
	// First argument should be a value of -128 to 127 range
	//
	// Second argument: 10 means decimal number
	// Third argument : 8 means 8-bits (int8)
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)

	// Now the val variable is int64 because ParseInt
	// returns an int64. But, since I passed 8 as its third
	// argument, it returns int8 convertable value.
	//
	// Try running the program with a value of -128 to 127
	// Running it beyond that range will result in
	// either -128 or 127.
	fmt.Println("int8 value is:", int8(val))

	// --------------------------------------
	// NOW IT'S YOUR TURN!
	// --------------------------------------

	// 1. Get an int16 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Printf("int16 value is: %d\n", val)

	// 2. Get an int32 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Printf("int32 value is: %d\n", val)

	// 3. Get an int64 value using ParseInt and convert it and print it
	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Printf("int64 value is: %d\n", val)

	// 4. Get an int8 value using ParseInt and convert it and print it
	//    But this time, get the value in bits.
	val, _ = strconv.ParseInt(os.Args[5], 2, 64)
	fmt.Printf("00000100 value is: %d\n", val)

	//    For example : 00000100
	//    Should print: 4
}
