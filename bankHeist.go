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
	fmt.Println(isHeistOn)
}