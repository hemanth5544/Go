package main

import (
	"fmt"
)
func main() {	
	fmt.Println("if else jinxx")


	loginCount := 5

	var result string

	if loginCount < 10 {

		result= "you are a new valid user"
	}else if loginCount >20 {
		result= "you are a valid user"
	}else {
		result= "you are a super user"
	}

	fmt.Println("Result:", result)

}