package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Join([]string{"Hello", "Go"}, " ")
	fmt.Printf("%s\n", result)

	parts := strings.Split("apple, orange, banana", ",")
	fmt.Printf("%s\n", parts)

	contains := strings.Contains("Gopher", "Go")
	fmt.Printf("%t\n", contains)

	count := strings.Count("bandung", "dung")
	fmt.Printf("%d\n", count)

	replaced := strings.Replace("Hello World", "World", "Go", 1)
	fmt.Printf("%s\n", replaced)

	lower := strings.ToLower("HELLO")
	upper := strings.ToUpper("world")

	fmt.Printf("%s\n", lower)
	fmt.Printf("%s\n", upper)

	trimmed := strings.TrimSpace("  Trim me  ")
	fmt.Printf("%s\n", trimmed)

	repeated := strings.Repeat("Go! ", 3)
	fmt.Printf("%s\n", repeated)

	compare := strings.Compare("apple", "banana")
	fmt.Printf("%d\n", compare)
}
