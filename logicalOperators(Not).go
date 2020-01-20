// In main.go add a ! operator to the if
//  condition so that "What are we waiting for??" 
//  is printed to the terminal.

package main

import "fmt"

func main() {
	readyToGo := true

	if !readyToGo {
		fmt.Println("Start the car!")
	} else {
		fmt.Println("What are we waiting for??")
	}
}
