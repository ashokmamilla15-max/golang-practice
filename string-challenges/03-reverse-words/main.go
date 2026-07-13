package main

import (
	"fmt"
	"strings"
)

func revrseWords(s string) string {
	if s == "" {
		return ""
	}
	words := strings.Split(s, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	input := "Hello, World!"
	reversedWords := revrseWords(input)
	fmt.Println("Reversed words:", reversedWords)
}
