package main

func main() {
	input := "Hello, World!"
	length := findLength(input)
	println("Length of the string:", length)
}

func findLength(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	for range s {
		count++
	}
	return count
}
