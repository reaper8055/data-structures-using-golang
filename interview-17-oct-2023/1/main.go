package main

import (
	"fmt"
	"sync"
)

func PrintOdd(wg *sync.WaitGroup, evenCh chan int) {
	defer wg.Done()
	for i := 1; i < 100; i += 2 {
		fmt.Println(i)
		fmt.Println(<-evenCh)
	}
}

func PrintEven(wg *sync.WaitGroup, evenCh chan int) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		evenCh <- i
		// fmt.Println(i)
	}
	close(evenCh)
}

func main() {
	var wg sync.WaitGroup
	evenCh := make(chan int)
	wg.Add(2)
	go PrintOdd(&wg, evenCh)
	go PrintEven(&wg, evenCh)
	wg.Wait()
}
