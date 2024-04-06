package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating of this app: ")
	userInput, _ := reader.ReadString('\n')
	fmt.Println("you have give a rating of ", userInput)
	fmt.Printf("the type of user input is: %T", userInput)
}