package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	numGoroutines := 6

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			data[i] = i * i
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("data: ", data)
}
