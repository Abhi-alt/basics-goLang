package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion app")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating: ")
	input, _ := reader.ReadString('\n')
	convertedInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	if err != nil {
		fmt.Println("please input a valid rating: ", err)
	} else {
		fmt.Println("we have added 1 to your rating ", convertedInput+1)
	}
}