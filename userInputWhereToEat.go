package main

import "fmt"

func main() {
	fmt.Println("What would you like to eat for lunch?")

	fmt.Scan(&food)

	fmt.Printf("Sure, we can have %v for lunch.", food)

}