package main

import (
	"fmt"
)

func FirstRune(str string) rune {
	runes := []rune(str)

	return runes[0]
}

func main() {
	str := "Ebrahim"

	fmt.Println(string(FirstRune(str)))
}
