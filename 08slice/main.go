package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices")
	fmt.Println("First method for slice")
	mySlice1 := []string{}
	fmt.Printf("the type of mySlice1 is %T\n", mySlice1) 
	mySlice1 = append(mySlice1, "Cheese", "Bread", "Mushroom")
	fmt.Println(mySlice1)
	fmt.Println("=========================")
	fmt.Println("Second way to initialize slice")
	var mySlice2 = make([]int, 4) // make method initialize memory or resrve non zero memory
	mySlice2[0] = 12
	mySlice2[1] = 13
	mySlice2[2] = 14
	mySlice2[3] = 15
	// mySlice2[4] = 16 // if you try to do this you will get error 
	fmt.Println(mySlice2)
	// but append method will work as it will append value as initial was just default value
	mySlice2 = append(mySlice2, 21, 22, 23)
	fmt.Println(mySlice2)
	fmt.Println("=========================")
	fmt.Println("Make slice or part of slice")
	fmt.Println(mySlice2[3:5])
	fmt.Println(mySlice2[:5])
	fmt.Println(mySlice2[1:])
	fmt.Println("================================")
	fmt.Println("Delete value based on index from slice")
	var courses = []string{"ReactJS", "NodeJS", "GoLang", "Java", "Python"}
	var index = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}