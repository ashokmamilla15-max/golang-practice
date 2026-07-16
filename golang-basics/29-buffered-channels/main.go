// A buffered channel, created with make(chan T, capacity), lets sends
// succeed without a waiting receiver until the buffer is full. This
// differs from unbuffered channels, which always block until both sides
// are ready.
package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "first"
	ch <- "second" // does not block: buffer has room
	fmt.Println("buffered sends without a receiver, length:", len(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
