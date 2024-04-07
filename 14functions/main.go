package main

import "fmt"

func main(){
	fmt.Println("==========welcome to functions===========")
	val := adder(12, 3)
	fmt.Println("This is the value of adder function: ", val)
	sumVal, msg := proAdder(1, 2, 4, 5, 6)
	fmt.Println("The value of addition is: ", sumVal)
	fmt.Println("The returned msg is: ", msg)
}

func adder(inpOne int, inpTwo int) int {
	return inpOne+inpTwo
}

func proAdder(inps ...int) (int, string){
	total := 0
	for _, inp := range inps{
		total += inp
	}

	return total, "This is returning from proAdder"
}