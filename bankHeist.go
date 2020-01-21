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
		fmt.Println("Looks like You've managed to make it past the guards.  Good Job, but Remember,
			this is the first step.")
	}
}