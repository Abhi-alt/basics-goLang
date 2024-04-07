package main

import "fmt"

func main(){
	fmt.Println("Welcome to if else conditional statement")
	loginCount := 27
	var msg string
	if loginCount < 27{
		msg = "Standard user"
		fmt.Println(msg)
	} else if loginCount > 27 {
		msg = "Pro user!!!"
		fmt.Println(msg)
	} else {
		fmt.Println("Just met the target")
	}
}