package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())


	isHeistOn := true
	eludedGuards := rand.Intn(100)

	fmt.Println(eludedGuards)

	if (eludedGuards >= 50) {
		fmt.Println("Looks like You've managed to make it past the guards.  Good Job, but Remember, this is the first step.")
	} else {
		isHeistOn := false
		fmt.Println(isHeistOn, "Plan a better disguise")
	}

	openedVault := rand.Intn(100)

	if (isHeistOn == true && openedVault >= 70) {
		fmt.Println("Grand and GO!")
	}

	leftSafely := rand.Intn(5)

	if (isHeistOn == true) {
		switch leftSafely {
		case 0:
			isHeistOn := false
			fmt.Println("isHeistOn:", isHeistOn)
		case 1:
			isHeistOn := false
			fmt.Println(isHeistOn, "turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn := false
			fmt.Println(isHeistOn, "Cops busted you!")
		case 3: 
			isHeistOn := false
			fmt.Println(isHeistOn, "You were ratted out")
		case 4: 
			isHeistOn := false
			fmt.Println(isHeistOn, "4 means you lose")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

w}