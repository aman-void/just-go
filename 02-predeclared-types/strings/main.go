package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tag := "café"
	fmt.Println("bytes:", len(tag), "letters:", utf8.RuneCountInString(tag))

	for i, r := range tag {
		fmt.Printf("byte %d -> %c\n", i, r) // r is a rune (int32)
	}
}
