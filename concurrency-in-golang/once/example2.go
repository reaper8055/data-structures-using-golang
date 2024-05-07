package main

import (
	"fmt"
	"sync"
)

// sync.Once only counts the number of times `Do` is called,
// not how many times unique functions passed into `Do` are
// called.
func main() {
	var once sync.Once
	var count int

	increment := func() { count++ }
	decrement := func() { count-- }

	once.Do(increment)
	once.Do(decrement)
	fmt.Println("count in example2: ", count)
}
