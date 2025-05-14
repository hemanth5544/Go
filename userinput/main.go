package main

import  (
	"fmt"
	"bufio"
	"os"

)

func main ()  {


	welcome := "Welcome to the Go programming language!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")

	//comma ok || err ok

	input,_ := reader.ReadString('\n')
	fmt.Println("Hello, " + input + "!")
	
}
