package main

import "fmt"

func main() {
	name := "H. J. Simp"

	// Add a switch statement that takes the condition name.
	switch name {
	// Add one case statement with the value "Butch" and inside its block print the string "Head to Robbers Roost.".
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	// Add a second case statement with the value "bonnie" and inside its block, print the string "Stay put in Joplin.".
	case "bonnie":
		fmt.Println("Stay put in Joplin.")
	// Add a default statement that will print out "Just hide!".
	default:
		fmt.Println("Just hide!")
	}
}