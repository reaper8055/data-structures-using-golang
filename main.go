package main

import (
	"fmt"

	slice "github.com/reaper8055/data-structures-using-golang/slices"
)

func main() {
	a := []int{1, 2, 3, 4}
	s := slice.New(a)

	b, _ := s.ValueAt(2)
	fmt.Println(b)

}
