// A goroutine is a lightweight, concurrently executing function started
// with the `go` keyword. Goroutines run independently of the caller, so
// synchronization (e.g. a WaitGroup or channel) is needed to wait for them
// to finish.
package main

import (
	"fmt"
	"sync"
)

func sayHello(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello from goroutine", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go sayHello(i, &wg)
	}
	wg.Wait()
	fmt.Println("all goroutines finished")
}
