package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerWithCancel(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: stop the context.WithCancel\n", name)
			return
		default:
			fmt.Printf("%s: works\n", name)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func workerWithTimeout(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: stop the WithTimeout\n", name)
			return
		default:
			fmt.Printf("%s: works\n", name)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func workerWithChannel(stopChan <-chan struct{}, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Printf("%s: stop the WithChannel\n", name)
			return
		default:
			fmt.Printf("%s: works\n", name)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func workerWithFlag(stopFlag *bool, wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for {
		if *stopFlag {
			fmt.Printf("%s: stop the WithFlag\n", name)
			return
		}
		fmt.Printf("%s: works\n", name)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup

	ctxCancel, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go workerWithCancel(ctxCancel, &wg, "workerWithCancel")

	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout()
	wg.Add(1)
	go workerWithTimeout(ctxTimeout, &wg, "workerWithTimeout")

	stopChan := make(chan struct{})
	wg.Add(1)
	go workerWithChannel(stopChan, &wg, "workerWithChannel")

	stopFlag := false
	wg.Add(1)
	go workerWithFlag(&stopFlag, &wg, "workerWithFlag")

	time.Sleep(1 * time.Second)

	cancel()

	time.Sleep(3 * time.Second)

	close(stopChan)

	stopFlag = true

	wg.Wait()
	fmt.Println("all goroutines are closed")

}
