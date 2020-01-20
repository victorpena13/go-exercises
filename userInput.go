package main

import "fmt"

func main() {
	fmt.Println("How are you doing?")

	var response string
	fmt.Scan(&response)

	fmt.Println("I'm %v.", response)
	//result if user types "good"
	//I'm good
	//result if user types "not bad"
	//I'm not
}

