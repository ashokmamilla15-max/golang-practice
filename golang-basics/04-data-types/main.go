// Go is statically typed. Basic types include integers (int, int8..int64,
// uint variants), floating point (float32, float64), booleans (bool),
// strings (string), and complex numbers (complex64, complex128).
package main

import "fmt"

func main() {
	var i int = 42
	var i8 int8 = 127
	var u uint = 42
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265358979
	var b bool = true
	var s string = "Go"
	var c complex128 = complex(2, 3)
	var by byte = 'A' // alias for uint8
	var r rune = 'A'  // alias for int32, represents a Unicode code point

	fmt.Println(i, i8, u, f32, f64, b, s, c, by, r)
	fmt.Printf("%T %T %T %T %T %T %T %T %T %T\n", i, i8, u, f32, f64, b, s, c, by, r)
}
