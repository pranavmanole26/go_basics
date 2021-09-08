package main

import "fmt"

// Type of enum is Month
type Month int

// Adding values in Month enum
const (
	Jan Month = iota + 1
	Feb
	Mar
	Apr
	May
	Jun
	Jul
	Aug
	Sep
	Oct
	Nov
	Dec
)

// Function which returns string value of enum values
func (m Month) getVal() string {
	return []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}[m-1]
}

func main() {
	fmt.Printf("Month number is %d and month is %s", Jan, Jan.getVal())
}
