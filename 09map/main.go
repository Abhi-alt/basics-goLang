package main

import "fmt"

func main(){
	fmt.Println("Welcome to map which is like object")
	var languages = make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Rubby"
	languages["GO"] = "GoLang"
	fmt.Printf("Type of the languages is: %T\n", languages)
	fmt.Println(languages)
	println(languages["JS"])
	fmt.Println("=================================")
	fmt.Println("Deleting a key pair value")
	delete(languages, "RB")
	fmt.Println(languages)
	fmt.Println("======Using for lop for value n key======")
	for key, value := range languages{
		fmt.Printf("For the key of %v : the value is %v\n", key, value)
	}
}