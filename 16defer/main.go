package main

import "fmt"

func main(){
	defer fmt.Println("World") // follows first in last out principle
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}