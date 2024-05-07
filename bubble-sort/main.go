package main

import "fmt"

func main() {
	s := []int{1, 10, 5, 0, 100, 99, 201, 400, 39, 33}
	fmt.Println(s)
	fmt.Println(bubbleSort(s))
}

func bubbleSort(s []int) []int {
	l := len(s) - 1
	for l >= 1 {
		for i := 0; i < l; i++ {
			if s[i] > s[i+1] {
				tmp := s[i]
				s[i] = s[i+1]
				s[i+1] = tmp
			}
		}
		l--
	}
	return s
}
