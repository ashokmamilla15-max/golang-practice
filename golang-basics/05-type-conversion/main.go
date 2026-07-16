// Go does not perform implicit type conversion between different types;
// conversions must be explicit using T(v).
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)

	// Converting between strings and numbers requires the strconv package,
	// not a plain type conversion.
	s := strconv.Itoa(i)
	n, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println("conversion error:", err)
	}
	fmt.Println(s, n)
}
