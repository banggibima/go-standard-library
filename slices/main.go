package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Bima", "Aryn", "Bayu", "Brian"}
	scores := []int{100, 95, 90, 85}

	minValue := slices.Min(scores)
	maxValue := slices.Max(scores)
	containsBima := slices.Contains(names, "Bima")
	indexAryn := slices.Index(names, "Aryn")

	fmt.Printf("%d\n", minValue)
	fmt.Printf("%d\n", maxValue)
	fmt.Printf("%v\n", containsBima)
	fmt.Printf("%d\n", indexAryn)
}
