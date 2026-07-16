// `defer` schedules a function call to run when the surrounding function
// returns, regardless of how it returns. Deferred calls run in LIFO order,
// and their arguments are evaluated immediately even though the call
// happens later.
package main

import "fmt"

func process() {
	fmt.Println("start processing")
	defer fmt.Println("deferred: cleanup 1")
	defer fmt.Println("deferred: cleanup 2")
	fmt.Println("end processing")
}

func main() {
	process()

	for i := 0; i < 3; i++ {
		defer fmt.Println("deferred in loop:", i)
	}
	fmt.Println("main body finished")
}
