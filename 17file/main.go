package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Welcome to create and read file")
	file, err1 := os.Create("./myTextFile.txt")
	checkError(err1)
	length, err2 := file.WriteString("This is awesome file created by goLang!!")
	checkError(err2)
	fmt.Println("The length of the file is: ", length)
	defer file.Close()
	content, err3 := os.ReadFile("./myTextFile.txt")
	checkError(err3)
	fmt.Println(content)
	fmt.Println(string(content))
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}