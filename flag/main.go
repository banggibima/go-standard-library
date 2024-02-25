package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var age int
	var married bool

	flag.StringVar(&name, "name", "Bayu", "User's name")
	flag.IntVar(&age, "age", 25, "User's age")
	flag.BoolVar(&married, "married", false, "Is the user married")

	flag.Parse()

	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%t\n", married)

	otherArgs := flag.Args()
	fmt.Printf("%s\n", otherArgs)
}
