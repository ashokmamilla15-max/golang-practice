// Strings in Go are immutable, UTF-8 encoded byte sequences. Indexing a
// string yields a byte, not a character; iterating with `range` yields
// runes (Unicode code points) and correctly decodes multi-byte characters.
package main

import "fmt"

func main() {
	s := "héllo"

	fmt.Println("byte length:", len(s))

	for i, r := range s {
		fmt.Printf("index %d: rune %q (%d)\n", i, r, r)
	}

	runes := []rune(s)
	fmt.Println("rune count:", len(runes))

	bytes := []byte(s)
	fmt.Println("as bytes:", bytes)

	fmt.Println("concatenation:", s+" world")
}
