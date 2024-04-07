package main

import "fmt"

func main() {
	fmt.Println("welcome to array cannot do much with array")
	var myArray [3]string //always need to give size of array while initializing it. You can define it empty but then you wont be able to insert anything to it
	myArray[0] = "Hello"
	myArray[1] = "World"
	fmt.Println(myArray, len(myArray)) // you will get empty space for not assigned space

	fruitList := [3]string {"Apple", "Orange", "Kiwi"}
	fmt.Println(fruitList, len(fruitList))
}