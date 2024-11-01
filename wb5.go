package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	duration := 5 * time.Second

	dataChannel := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataChannel)
				return
			case dataChannel <- rand.Intn(10):
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
	for data := range dataChannel {
		fmt.Printf("value received: %d\n", data)
	}
	fmt.Println("end")

}
