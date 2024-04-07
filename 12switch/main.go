package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println("Welcome to swicth case")
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber);
	switch diceNumber {
	case 1:
		fmt.Printf("You have got %v, try again to win", diceNumber)
	case 2:
		fmt.Printf("You have got %v, try again to win", diceNumber)
	case 3:
		fmt.Printf("Your value is %v and you have won", diceNumber)
	case 4:
		fmt.Printf("Your no value is %v, try again to win", diceNumber)
	case 5:
		fmt.Printf("Your no is %v and you won", diceNumber)
	case 6:
		fmt.Printf("Your no is %v and you got jackpot", diceNumber)
	default:
		fmt.Println("Thats not a correct no")
	}
}