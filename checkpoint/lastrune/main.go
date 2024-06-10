package main

import (
	"github.com/01-edu/z01"
)

func LastRune(s string) rune {
	str := []rune(s)
	lastRune := len(str) - 1

	return str[lastRune]
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ebrahim"))
	z01.PrintRune('\n')
}
