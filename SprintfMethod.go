// Assign to wish the value of calling fmt.Sprintf() with the values template and pet.

// wish should then contain the interpolated string "I wish I had a puppy.".

package main

import "fmt"

func main() {
	correctAns := "A"
	answer := fmt.Sprintf("And the correct answer is... %v!", correctAns)

	fmt.Print(answer)
	// Prints: And the correct answer is... A!

	template := "I wish I had a %v."
 	pet := "puppy"
  
  	var wish string
  
  	// Add your code below:

  	wish = fmt.Sprintf(template, pet)
  
  
  fmt.Println(wish)
}