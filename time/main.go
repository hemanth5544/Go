package main

import (
	"fmt"
	"time"
)


func main(){
	fmt.Println("Time jinxx")

//pointers
myNumber:= 10
var ptr =&myNumber
fmt.Println("Pointer to myNumber:", ptr) //address of myNumber
fmt.Println("Value of myNumber:", *ptr) //value in myNumber at the ptr

*ptr = *ptr+20
fmt.Println("Value of myNumber after adding 20:", myNumber) //value in myNumber at the ptr

	presentTIme := time.Now()
	fmt.Println("Current time:", presentTIme)
	fmt.Println("Current time iin format:", presentTIme.Format("01-02-2006 Monday"))
}






