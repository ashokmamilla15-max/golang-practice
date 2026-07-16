// Channels let goroutines communicate and synchronize by sending and
// receiving typed values with `<-`. An unbuffered channel blocks the
// sender until a receiver is ready, and vice versa.
package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
	}
	close(results)
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	go worker(jobs, results)

	for r := range results {
		fmt.Println("result:", r)
	}
}
