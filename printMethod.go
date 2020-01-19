package main

import "fmt"

func main() {
  fmt.Println("Let's first see how", "the Println() method works.")
  fmt.Println("Notice that each statement adds a newline for us.")
  fmt.Println("There's also a default space", "between the string arguments.")
  // result
  // Let's first see how the Println() method works.
  // Notice that each statement adds a newline for us.
  // There's also a default space between the 
 
  fmt.Print("Print", "is", "different")
  fmt.Print("see")
  //result
  //Printisdifferentsee


  //summary:
  //Println will format the statements with spaces and linebreaks
  //print does not format the output to include spaces or linebreaks 


}