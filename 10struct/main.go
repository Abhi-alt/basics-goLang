package main

import "fmt"

func main(){
	fmt.Println("Welcome to struct")
	var abhi = User{"Abhishek", 28, "Pune"}
	fmt.Printf("Type of abhi is %T\n", abhi)
	fmt.Printf("The value of abhi is %+v\n", abhi)
	fmt.Printf("Hi I am %v and I am %v years old and I work in %v\n", abhi.Name, abhi.age, abhi.city)
}

type User struct {
	Name string
	age int
	city string
}