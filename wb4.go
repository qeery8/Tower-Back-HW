package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	numWorkers := 3

	dataChannel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("program closing")
		cancel()
		close(dataChannel)
	}()

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dataChannel <- rand.Intn(10):
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, dataChannel, &wg)
	}

	wg.Wait()
}

func worker(ctx context.Context, id int, dataChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d closed work\n", id)
			return
		case data, ok := <-dataChannel:
			if !ok {
				return
			}
			fmt.Printf("worker %d get: %d\n", id, data)
		}
	}
}
