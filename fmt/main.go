package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, Go!\n")
	fmt.Println("Go!")

	name := "Bima"
	age := 30

	fmt.Printf("%s %d\n", name, age)
	fmt.Println(name, age)

	message := fmt.Sprintf("Hello, %s!", "World")
	fmt.Println(message)

	number := 42.6789

	fmt.Printf("%f\n", number)
	fmt.Printf("%.2f\n", number)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

	var str string
	fmt.Print("input string: ")
	fmt.Scanln(&str)
	fmt.Printf("sample output: %s\n", str)

	var num int
	fmt.Print("input integer: ")
	fmt.Scanln(&num)
	fmt.Printf("sample output: %d\n", num)
}
