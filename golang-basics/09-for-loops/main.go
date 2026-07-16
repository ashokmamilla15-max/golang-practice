// `for` is the only looping construct in Go; it can act as a classic
// counted loop, a while loop, an infinite loop, or a range loop over
// arrays, slices, strings, maps, and channels.
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("counted:", i)
	}

	n := 0
	for n < 3 {
		fmt.Println("while-style:", n)
		n++
	}

	for i, v := range []string{"a", "b", "c"} {
		fmt.Println("range:", i, v)
	}

	count := 0
	for {
		if count == 2 {
			break
		}
		fmt.Println("infinite loop iteration:", count)
		count++
	}
}
