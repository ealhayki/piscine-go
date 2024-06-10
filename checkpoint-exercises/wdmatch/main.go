package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ContainRunes(s1, s2 string) bool {
	count := make(map[rune]int)

	for _, ch := range s2 {
		count[ch]++
	}
	for _, ch := range s1 {
		if len(s1) == 0 {
			return false
		}
		count[ch]--
	}

	return true
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]

	if ContainRunes(first, second) {
		for _, ch := range first {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
