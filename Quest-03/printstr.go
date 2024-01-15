package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	str := []rune(s)

	for _, char := range str {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}

func main() {
	PrintStr("Hello World!")
}
