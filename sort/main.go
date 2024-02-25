package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	intSlice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sort.Ints(intSlice)
	fmt.Printf("%v\n", intSlice)

	isSorted := sort.IntsAreSorted(intSlice)
	fmt.Printf("%v\n", isSorted)

	strSlice := []string{"apple", "banana", "orange", "grape"}
	sort.Strings(strSlice)
	fmt.Printf("%v\n", strSlice)

	searchValue := "orange"
	index := sort.SearchStrings(strSlice, searchValue)
	fmt.Printf("%s: %d\n", searchValue, index)

	people := []Person{
		{"Bayu", 25},
		{"Aryn", 30},
		{"Brian", 20},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Printf("%v\n", people)
}
