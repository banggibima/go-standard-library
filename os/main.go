package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("%s\n", args)

	envVar := os.Getenv("HOME")
	fmt.Printf("%s\n", envVar)

	wd, _ := os.Getwd()
	fmt.Printf("%s\n", wd)

	os.Mkdir("new_directory", 0755)

	file, _ := os.Create("new_file.txt")
	defer file.Close()

	file.WriteString("Hello, Go!")

	content, _ := os.ReadFile("new_file.txt")
	fmt.Printf("%s\n", string(content))

	os.Remove("new_file.txt")
	os.Remove("new_directory")
}
