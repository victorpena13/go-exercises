package main

import "fmt"

func main() {
	fmt.Println("How have you been?")

	var response1 string
	var response2 string
	fmt.Scan(&response1)
	fmt.Scan(&response2)

	fmt.Println("You know, I am %v %v", response1, response2)
}

