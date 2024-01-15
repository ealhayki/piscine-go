package main

import "fmt"

func StrLen(s string) int {
	str := []rune(s)
	strLen := len(str)

	return strLen
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
