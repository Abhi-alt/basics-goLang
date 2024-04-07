package main

import (
	"fmt"
	"io"
	"net/http"
)

const uri = "https://jsonplaceholder.typicode.com/users"

func main(){
	fmt.Println("Welcome to http request making")
	resp, err := http.Get(uri)
	checkError(err)
	defer resp.Body.Close()
	dataBytes, err := io.ReadAll(resp.Body)
	checkError(err)
	fmt.Println(string(dataBytes))
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}