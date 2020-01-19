package main

import "fmt"

func main() {
specialNum := 42
fmt.Printf("This value's type is %T.", specialNum)
//result: This value's type is int

fmt.Println("\n***") 
// Added for spacing

quote := "To do or not to do"
fmt.Printf("This value's type is %T.", quote)
//Result: This value's type is string

fmt.Println("\n***") 
// Added for spacing

votingAge := 21
fmt.Printf("You must be %d years old to vote.", votingAge)
//Result: You must be 21 years old to vote.

fmt.Println("\n***") 
// Added for spacing

gpa := 3.8
fmt.Printf("You're averaging: %f.", gpa)
//Result: You're averaging 3.800000.


fmt.Println("\n***") 
// Added for spacing

gpa := 3.8
fmt.Printf("You're averaging: %.2f.", gpa)
//Result: You're averaging 3.80


fmt.Println("\n***") 
// Added for spacing

  floatExample := 1.75
  fmt.Printf("Working with a %T", floatExample)
  // prints: Working with a float64
  
  fmt.Println("\n***") // Added for spacing
  
  yearsOfExp := 3
  reqYearsExp := 15
  fmt.Printf("I have %d years of Go experience and this job is asking for %d years.", yearsOfExp, reqYearsExp) 
  // prints: I have 3 years of Go experience and this job is asking for 15 years.
  
  fmt.Println("\n***") // Added for spacing
  
  stockPrice := 3.50
  fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice) 
  // prints: Each share of Gopher feed is $3.50!
}