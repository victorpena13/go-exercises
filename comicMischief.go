package main

import "fmt"

func main() {
  var publisher string
  var writer string
  var artist string
  var title string
  var year int
  var pageNumber int
  var grade float32
  var condition float32
  
  title = "Mr.GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  condition = 6.5
  
  fmt.Println(title, "written by", writer, "drawn by", artist, "year:", year, "grade:", grade, "page number:", pageNumber, "publisher", publisher, "condition:",condition)
  
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0
  
  fmt.Println(title, "written by", writer, "drawn by", artist, "year:", year, "grade:", grade, "page number:", pageNumber, "publisher", publisher, "condition:",condition)
  
}