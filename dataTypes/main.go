package main

import (
	"fmt"

)
func main() {
	fmt.Println("Arrays jinxx")

	var myArray [4]int;
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40

	fmt.Println("Array:", myArray)

	var initializedArray = [4]int{10, 20, 30, 40}
	fmt.Println("Initialized Array:", initializedArray)


language:= make(map[string]string)
language["en"] = "English"
language["fr"] = "French"
language["es"] = "Spanish"	

fmt.Println("Map:", language)


fmt.Println("Map of en:", language["en"])


delete(language, "en")
fmt.Println("Map after deletion:", language)


for key,value:= range language {
	fmt.Println("Key:", key, "Value:", value)



}
}