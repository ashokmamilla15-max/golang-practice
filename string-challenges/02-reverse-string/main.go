package main

func main() {
	input := "Hello, World!"
	reversed := reverseString(input)
	println("Reversed string:", reversed)

}
func reverseString(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
