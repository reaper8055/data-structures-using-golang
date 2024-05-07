package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "correlate"
	reverse(s)

}

type foundVowels struct {
	ch  string
	pos int
}

func reverse(s string) string {

	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	// correlate —> cerraleto

	// c o r r e l a t e
	// 1. split (slice)
	// 2. position of all vowels
	// 3. reverse all vowels
	// 4. put them back into the word

	splitWord := strings.Split(s, "")

	var found []foundVowels
	var vs string

	for pos, ch := range splitWord {
		if vowels[ch] {
			found = append(found, foundVowels{
				ch:  ch,
				pos: pos,
			})
			vs = ch + vs
		}
	}

	vss := strings.Split(vs, "")

	for pos := range splitWord {
		for i, f := range found {
			if pos == f.pos {
				splitWord[pos] = vss[i]
			}
		}
	}
	fmt.Println(splitWord)

	return ""
}
