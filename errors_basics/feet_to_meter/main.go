package main

import (
	"fmt"
	"os"
	"strconv"
)

func feetToMeter(lf float64) float64 {
	return lf * 0.3048
}

func inputLength() (lf float64) {
	fmt.Println("Enter length in feet:")
	fmt.Scanf("%f", &lf)
	return
}

func main() {

	var lenFt, lenMt float64
	var err error
	if len(os.Args) != 1 {
		lenFt, err = strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			lenFt = inputLength()
		}
	} else {
		lenFt = inputLength()
	}
	lenMt = feetToMeter(lenFt)
	fmt.Printf("Length %0.2f feet is converted to %0.2f meter.\n", lenFt, lenMt)
}
