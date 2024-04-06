package main

import "fmt"

func main() {
	var username string = "abhi01"
	fmt.Println(username)
	fmt.Printf("The type of this variable is: %T \n", username)

	var count uint8 = 255
	fmt.Println(count)
	fmt.Printf("The type of this variable is: %T \n", count)

	var count16 uint16 = 258
	fmt.Println(count16)
	fmt.Printf("The type of this variable is: %T \n", count16)

	var countFloat float32 = 12.34456667788643222
	fmt.Println(countFloat)
	fmt.Printf("The type of this variable is: %T \n", countFloat)

	var countFloat64 float64 = 12.34456667788643222
	fmt.Println(countFloat64)
	fmt.Printf("The type of this variable is: %T \n", countFloat64)

	var countInt int = 45434
	fmt.Println(countInt)
	fmt.Printf("the type of this variable is %T \n", countInt)

}