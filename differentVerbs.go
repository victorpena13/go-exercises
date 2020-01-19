package main

import "fmt"

func main() {
specialNum := 42
fmt.Printf("This value's type is %T.", specialNum)
//result: This value's type is int

quote := "To do or not to do"
fmt.Printf("This value's type is %T.", quote)
//Result: This value's type is string

votingAge := 21
fmt.Printf("You must be %d years old to vote.", votingAge)
//Result: You must be 21 years old to vote.

gpa := 3.8
fmt.Printf("You're averaging: %f.", gpa)
//Result: You're averaging 3.800000.

gpa := 3.8
fmt.Printf("You're averaging: %.2f.", gpa)
//Result: You're averaging 3.80
}