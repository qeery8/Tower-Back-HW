package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	sumSquares := 0

	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			square := n * n
			mu.Lock()
			sumSquares += square
			mu.Unlock()
		}(num)
	}

	wg.Wait()
	fmt.Printf("Summary square = %d\n", sumSquares)
}
