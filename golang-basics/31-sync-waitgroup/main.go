// sync.WaitGroup coordinates a fixed number of goroutines: call Add before
// starting each one, Done when a goroutine finishes, and Wait to block
// until all of them have called Done.
package main

import (
	"fmt"
	"sync"
)

func fetch(id int, results *[]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	value := id * 10
	mu.Lock()
	*results = append(*results, value)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]int, 0)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go fetch(i, &results, &mu, &wg)
	}
	wg.Wait()

	fmt.Println("collected", len(results), "results")
}
