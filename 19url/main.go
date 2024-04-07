package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://abhi.go.dev/role?id=1234&name=abhi&city=pune"

func main(){
	fmt.Println("Welcome to url world")
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawPath)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Scheme)

	qparams := result.Query()
	for key, val := range qparams {
		fmt.Println(key, " : ", val)
	}
	mynewurl := &url.URL{
		Host: "abhi.go",
		Path: "/track",
		RawQuery: "id=111&place=mh",
		Scheme: "https",
	}
	fmt.Println(mynewurl.String())
}