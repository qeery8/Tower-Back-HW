package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- num * num
}

func main() {
	num := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(num))
	var wg sync.WaitGroup

	for _, n := range num {
		wg.Add(1)
		go square(n, &wg, ch)
	}
	wg.Wait()
	close(ch)

	for result := range ch {
		fmt.Println(result)
	}

}
