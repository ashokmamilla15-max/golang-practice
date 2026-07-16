// Go's switch does not fall through by default (unlike C/Java). It also
// supports switching on no expression (acting like if/else if chains) and
// explicit fallthrough.
package main

import "fmt"

func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "weekend"
	default:
		return "weekday"
	}
}

func main() {
	fmt.Println(dayType("Sunday"), dayType("Monday"))

	n := 15
	switch {
	case n%15 == 0:
		fmt.Println("FizzBuzz")
	case n%3 == 0:
		fmt.Println("Fizz")
	case n%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(n)
	}

	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three (via fallthrough)")
	}
}
