package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{"test"}

func main() {
	fmt.Println("Go routines working")
	websites := []string{
		"https://fb.com",
		"https://google.com",
		"https://github.com",
		"https://zomato.com",
		"https://x.com",
	}
	for _, w := range websites {
		go getStatus(w)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
	// go greeter("hello")
	// greeter("world")
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatus(website string) {
	defer wg.Done()
	resp, err := http.Get(website)
	if err != nil {
		fmt.Println("OOPS seems website is down")
	} else {
		mut.Lock()
		signals = append(signals, website)
		mut.Unlock()
		fmt.Printf("%d is status code for url %s\n", resp.StatusCode, website)
	}
}
