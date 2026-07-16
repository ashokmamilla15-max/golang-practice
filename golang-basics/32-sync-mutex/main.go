// sync.Mutex provides mutual exclusion: Lock/Unlock ensures only one
// goroutine at a time accesses a shared resource, preventing data races
// like lost updates on a shared counter.
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	fmt.Println("final counter value:", counter.value)
}
