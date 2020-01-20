package main

import "fmt"

func main() {
	rightTime := true
	rightPlace := true

	// Edit this condition for the FIRST checkpoint
	// In the first conditional provided, 
	// in addition to checking rightTime, 
	// use the && operator to check if 
	// rightPlace is also true.
	if rightTime {
		fmt.Println("We're outta here!")
	} else {
		fmt.Println("Be patient...")
	}


	// In the second conditional provided, 
	// use the || operator to check for enoughBags.
	enoughRobbers := false
	enoughBags := true

	if enoughRobbers || enoughBags {
		fmt.Println("Grab everything!")
	} else {
		fmt.Println("Grab whatever you can!")
	}
}
