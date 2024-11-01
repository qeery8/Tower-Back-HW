package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	in := make(chan int)
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			in <- n
		}
		close(in)
	}()

	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()

	for result := range out {
		fmt.Println(result)
	}
}
