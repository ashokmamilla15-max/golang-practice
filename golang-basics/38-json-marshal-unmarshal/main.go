// The encoding/json package converts Go values to and from JSON. Struct
// tags (e.g. `json:"name"`) control field names, and json.Marshal /
// json.Unmarshal do the conversion.
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println("JSON:", string(data))

	var decoded Person
	jsonInput := `{"name":"Bob","age":25,"email":"bob@example.com"}`
	if err := json.Unmarshal([]byte(jsonInput), &decoded); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Printf("decoded: %+v\n", decoded)
}
