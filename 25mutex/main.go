package main

import (
	"fmt"
	"sync"
)

func main() {
	mut := &sync.RWMutex{}
	wg := &sync.WaitGroup{}

	var scores = []int{1}

	wg.Add(3)
	func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("GR 1")
		m.Lock()
		scores = append(scores, 1)
		m.Unlock()
		m.RLock()
		fmt.Println(scores)
		m.RUnlock()
		wg.Done()
	}(wg, mut)
	func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("GR 2")
		m.Lock()
		scores = append(scores, 2)
		m.Unlock()
		m.RLock()
		fmt.Println(scores)
		m.RUnlock()
		wg.Done()
	}(wg, mut)
	func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("GR 3")
		m.RLock()
		fmt.Println(scores)
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
}
