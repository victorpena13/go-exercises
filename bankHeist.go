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
		fmt.Println("Plan a better disguise")
		fmt.Println(isHeistOn)
	}

	openedVault := rand.Intn(100)

	if (isHeistOn == true && openedVault >= 70) {
		fmt.Println("Grand and GO!")
	}

	fmt.Println("isHeistOn is currently:", isHeistOn)
}