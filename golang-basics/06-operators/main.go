// Go supports arithmetic (+ - * / %), comparison (== != < > <= >=),
// logical (&& || !), bitwise (& | ^ &^ << >>), and assignment (=, +=, ...)
// operators.
package main

import "fmt"

func main() {
	a, b := 10, 3

	fmt.Println("Arithmetic:", a+b, a-b, a*b, a/b, a%b)
	fmt.Println("Comparison:", a == b, a != b, a > b, a < b)
	fmt.Println("Logical:", a > 5 && b < 5, a > 5 || b > 5, !(a > 5))
	fmt.Println("Bitwise AND/OR/XOR/AND-NOT:", a&b, a|b, a^b, a&^b)
	fmt.Println("Shift left/right:", a<<1, a>>1)

	a += 5
	fmt.Println("After a += 5:", a)
}
