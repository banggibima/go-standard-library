package main

import (
	"fmt"
	"math"
)

func main() {
	exp := math.Pow(2, 3)
	fmt.Printf("%v\n", exp)

	sqrt := math.Sqrt(25)
	fmt.Printf("%v\n", sqrt)

	log := math.Log(10)
	fmt.Printf("%v\n", log)

	sine := math.Sin(math.Pi / 2)
	fmt.Printf("%v\n", sine)

	cosine := math.Cos(math.Pi)
	fmt.Printf("%v\n", cosine)

	fmt.Printf("%v\n", math.Pi)
	fmt.Printf("%v\n", math.E)
}
