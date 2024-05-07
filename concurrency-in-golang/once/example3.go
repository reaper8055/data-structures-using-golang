package main

import (
	"sync"
)

// This program will deadlock because the call to `Do` at 1 won't proceed
// until the call to `Do` at 2 exits.
func main() {
	var onceA, onceB sync.Once
	var initB func()

	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // 1
	onceA.Do(initA)                    // 2
}
