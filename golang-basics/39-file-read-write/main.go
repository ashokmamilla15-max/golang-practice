// The os package handles files. os.WriteFile and os.ReadFile cover simple
// whole-file operations; os.CreateTemp is used here so the example is
// self-contained and cleans up after itself.
package main

import (
	"fmt"
	"os"
)

func main() {
	tmpFile, err := os.CreateTemp("", "golang-basics-*.txt")
	if err != nil {
		fmt.Println("create error:", err)
		return
	}
	path := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(path)

	content := []byte("Hello from Go file I/O!")
	if err := os.WriteFile(path, content, 0644); err != nil {
		fmt.Println("write error:", err)
		return
	}

	read, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println("file contents:", string(read))
}
