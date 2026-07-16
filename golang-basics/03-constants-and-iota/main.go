// Constants are declared with `const` and must be assignable at compile
// time; they cannot use `:=` or be reassigned. `iota` is a special
// identifier that resets to 0 at each `const` block and increments by one
// per line, commonly used to build enumerations.
package main

import "fmt"

const Pi = 3.14159

const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

const (
	_  = iota // skip 0
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println("Pi:", Pi)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("KB:", KB, "MB:", MB, "GB:", GB)
}
