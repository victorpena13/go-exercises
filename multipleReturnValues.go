package main 

import "fmt"

func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
	averageGrade := (midtermGrade + finalGrade) / 2
	var gradeLetter string
	if averageGrade > 90 {
		gradeLetter = "A"
	} else if average > 80 {
		gradeLetter = "B"
	} else if average > 70 {
		gradeLetter = "C"
	} else if average > 60 {
		gradeLetter = "D"
	} else {
		gradeLetter = "F"
	}

	return gradeLetter, averageGrade
}

func main() {
	var myMidterm, myFinal float32
	myMidterm = 89.4
	myFinal = 74.9
	var myAverage 
}