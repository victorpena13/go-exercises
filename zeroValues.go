package main

import "fmt"

func main() {

  var emptyInt int8
  var emptyFloat float32
  var emptyString string
  fmt.Println(emptyInt, emptyFloat, emptyString)
}

//since no value was assigned to variable 
//the variables meant to hold a number value will 
//result in 0 and the string will be blank.