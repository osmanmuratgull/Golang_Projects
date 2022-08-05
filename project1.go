package main

import (
	"fmt"
	"unicode/utf8"
)

func MirrorWriting(str string) string {
	byte_str := []rune(str)

	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {

		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]

	}

	return string(byte_str)

}

func main() {

	s := "Osman Muratgül"

	i := utf8.RuneCountInString(s)
	fmt.Println(s)
	fmt.Println(i)

	a := "Golang"

	l := utf8.RuneCountInString(a)
	fmt.Println(a)
	fmt.Println(l)

	fmt.Println(MirrorWriting("Osman Muratgül"))
	fmt.Println(MirrorWriting("Golang"))

}
