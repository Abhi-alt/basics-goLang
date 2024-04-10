package main

import (
	"fmt"
	"sync"
)

func main() {
	myChan := make(chan string, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	func(wg *sync.WaitGroup, c chan string) {
		c <- "hello from me"
		wg.Done()
	}(wg, myChan)
	func(wg *sync.WaitGroup, c chan string) {
		fmt.Println(<-c)
		wg.Done()
	}(wg, myChan)

	wg.Wait()
}
