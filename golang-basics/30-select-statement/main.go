// `select` waits on multiple channel operations at once and proceeds with
// whichever case is ready first, similar to a switch for channels. A
// `default` case makes the select non-blocking.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	empty := make(chan int)
	select {
	case v := <-empty:
		fmt.Println(v)
	default:
		fmt.Println("no value ready, default case")
	}
}
