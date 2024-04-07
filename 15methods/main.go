package main

import "fmt"

func main(){
	Abhi := User{"Abhishek", 28, "abhi@go.dev"}
	fmt.Printf("Hi %v you are %v yrs old and your email id is %v\n", Abhi.Name, Abhi.Age, Abhi.Email)
	Abhi.GreetUser()
	Abhi.NewMail()
	fmt.Printf("Hi %v you are %v yrs old and your email id is %v\n", Abhi.Name, Abhi.Age, Abhi.Email)
}

type User struct {
	Name string
	Age  int
	Email string 
}

//creating methods to struct

func (u User) GreetUser() {
	fmt.Println("Hello from GoLang Mr. ", u.Name)
}

func (u User) NewMail() {
	u.Email = "abhishek@go.dev"
	fmt.Println("Your email is ", u.Email, " but still original email wont change as its justa copy and not pointer")
}