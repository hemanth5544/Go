package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch jinxx")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number:", diceNumber)

	switch diceNumber {

	case 1:
		fmt.Println("Dice number is 1 you can open")
	case 2:
		fmt.Println("Dice number is 2 you can move 2 steps")

	case 3:

		fmt.Println("Dice number is 3 you can move 3 steps")
		fallthrough
	// fallthrough will execute the next case even if the condition is not met
	case 4:

		fmt.Println("Dice number is 4 you can move 4 steps")

	case 5:
		fmt.Println("Dice number is 5 you can move 5 steps")
	case 6:
		fmt.Println("Dice number is 6 you can move 6 steps")

	default:
		fmt.Println("Invalid dice number")

	}

}
