package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	joinedPath := filepath.Join("dir", "subdir", "file.txt")
	cleanedPath := filepath.Clean("/../dir/subdir/../file.txt")

	fmt.Printf("%s\n", joinedPath)
	fmt.Printf("%s\n", cleanedPath)

	dir, file := filepath.Split("/path/to/file.txt")
	fmt.Printf("%s\n", dir)
	fmt.Printf("%s\n", file)

	baseName := filepath.Base("/path/to/file.txt")
	extension := filepath.Ext("/path/to/file.txt")

	fmt.Printf("%s\n", baseName)
	fmt.Printf("%s\n", extension)

	isAbsolute := filepath.IsAbs("/path/to/file.txt")
	fmt.Printf("%t\n", isAbsolute)

	absPath, _ := filepath.Abs("../path/to/file.txt")
	fmt.Printf("%s\n", absPath)
}
