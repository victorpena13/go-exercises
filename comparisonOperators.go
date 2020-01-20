package main

import "fmt"

func main() {
	lockCombo := "2-35-19"
	robAttempt := "1-1-1"

	// create a if statement that checks if 
	// lockCombo and robAttempt are the same. 
	// If the condition evaluates to true, 
	// print out "The vault is now opened.".
	//Add your code below:
if (lockCombo == robAttempt) {
	fmt.Println("the vault is now open")
	} else {
		fmt.Println("the vault is locked")
	}
}