// Arrays have a fixed size that is part of their type ([N]T). Assigning
// or passing an array copies all of its elements.
package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[4] = 50

	fmt.Println(arr, len(arr))

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	inferred := [...]int{1, 2, 3}
	fmt.Println(inferred, len(inferred))

	copyArr := primes
	copyArr[0] = 999
	fmt.Println("original unaffected:", primes)
	fmt.Println("copy modified:", copyArr)
}
