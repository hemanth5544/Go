package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	Hemanth := User{"Hemanth", "rachap@gmail.com", true, 25}
	Hemanth.GetStatus()
	Hemanth.NewMail()


	fmt.Printf("Heamnth detils are %+v\n", Hemanth) // %v for value, %+v for key value pair

}

func (u User) GetStatus() {
	fmt.Println("User status is", u.Status)

}

func (u User) NewMail() string {
	u.Email = "new@gmail"
	fmt.Println("New email is", u.Email)
	return u.Email
}
