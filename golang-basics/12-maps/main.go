// Maps are Go's built-in hash table type, declared as map[KeyType]ValueType.
// Looking up a missing key returns the value type's zero value; the
// "comma ok" idiom distinguishes a missing key from a zero-valued one.
package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	ages["Carol"] = 28
	fmt.Println(ages)

	age, ok := ages["Dave"]
	fmt.Println("Dave's age:", age, "found:", ok)

	delete(ages, "Bob")
	fmt.Println("after delete:", ages)

	for name, a := range ages {
		fmt.Println(name, "is", a)
	}

	fmt.Println("map length:", len(ages))
}
