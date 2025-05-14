package main

import "fmt"

func main() {
	greeter()

	result := add(5, 10)
	fmt.Println("Result of addition:", result)

	superResult := superAdder(1, 2)
	fmt.Println("Result of super addition:", superResult)
}

func greeter() {

	fmt.Println("Hello from greeter")
}

func add(x int, y int) int {

	return x + y
}


func superAdder(vaues ...int)int{
	sum:=0;
	for _,val:=range vaues{
		sum+=val
	}
	return sum
}