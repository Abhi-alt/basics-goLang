package main

import "fmt"

func main(){
	fmt.Println("Welcome to pointer function")

	var ptr *int
	fmt.Println("this is empty pointer ", ptr)
    num := 9
	ptr = &num
	println(ptr)

	myNum := 23
	myNumPtr := &myNum //this is how you ref to the pointer of any variable
	fmt.Println(myNumPtr)
	fmt.Println(*myNumPtr)
	*myNumPtr = 2 * *myNumPtr
	fmt.Println(myNum, myNumPtr)
}