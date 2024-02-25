package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "Go"
	text := "Golang is awesome!"
	match := regexp.MustCompile(pattern).MatchString(text)
	fmt.Printf("%t\n", match)

	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString("There are 42 cats and 7 dogs.", -1)
	fmt.Printf("%s\n", numbers)

	replaced := re.ReplaceAllString("Price: $100. Quantity: 3.", "XX")
	fmt.Printf("%s\n", replaced)

	re = regexp.MustCompile(`(\w+)\s(\w+)`)
	submatches := re.FindStringSubmatch("John Doe, Jane Smith")
	fmt.Printf("%v %v\n", submatches[1], submatches[2])

	text = "apple orange banana"
	result := regexp.MustCompile(`\w+`).ReplaceAllStringFunc(text, func(s string) string {
		return s[:2]
	})
	fmt.Printf("%s\n", result)
}
