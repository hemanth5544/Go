package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your age: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Your age is:", input)

	// Convert to int
	numAge, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return
	}

	fmt.Println("Your age as an int:", numAge+1)
	fmt.Printf("Type of numAge: %T\n", numAge)
}
