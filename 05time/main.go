package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("this is my time module")
	presentTime := time.Now()
	fmt.Println(presentTime)
	// for formating time we need to use 01-02-2006 format only. refer docs for more format
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	createdTime := time.Date(1995, time.July, 27, 12, 12, 0, 0, time.Local)
	fmt.Println(createdTime)
	fmt.Println(createdTime.Format("01-02-2006 Monday"))

}