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

	Text1 := "Osman Muratgül"

	Text1Count := utf8.RuneCountInString(Text1)

	Text2 := "Golang"

	Text2Count := utf8.RuneCountInString(Text2)

	fmt.Println(MirrorWriting("Osman Muratgül"))
	fmt.Println(MirrorWriting("Golang"))
	fmt.Println(Text1Count)
	fmt.Println(Text2Count)

}
