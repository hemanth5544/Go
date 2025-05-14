package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs jinxx")

hemanht:= User{"Hemanth", "rachap@gmail.com", true, 25}
fmt.Printf("Heamnth detils are %+v\n", hemanht)// %v for value, %+v for key value pair
fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", hemanht.Name, hemanht.Email, hemanht.Status, hemanht.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
