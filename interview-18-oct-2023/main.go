package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	go send(ch)

	for j := 0; j < 20; j++ {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go worker(&wg, ch)
		}
	}
	wg.Wait()
}

// 200 tasks
// can only process 10 max at a time
func worker(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Println(<-ch)
}

func send(ch chan int) {
	for i := 0; i < 200; i++ {
		ch <- i
	}
	close(ch)
}
