package main

import (
	"errors"
	"fmt"
)

func multiply(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("anything multiplied by 0 is equal to 0")
	} else {
		return a * b, nil
	}
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("anything divided by 0 is undefined")
	} else {
		return float64(a) / float64(b), nil
	}
}

func main() {
	product, err := multiply(10, 2)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("%d\n", product)
	}

	product, err = multiply(8, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("%d\n", product)
	}

	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("%f\n", quotient)
	}

	quotient, err = divide(8, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("%f\n", quotient)
	}
}
