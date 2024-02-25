package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := strconv.Itoa(42)
	fmt.Printf("%s\n", numStr)

	num, _ := strconv.Atoi("123")
	fmt.Printf("%d\n", num)

	numBase, _ := strconv.ParseInt("1010", 2, 0)
	fmt.Printf("%d\n", numBase)

	numFormat := strconv.FormatInt(42, 16)
	fmt.Printf("%s\n", numFormat)

	floatNum, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("%f\n", floatNum)

	floatFormat := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("%s\n", floatFormat)

	boolVal, _ := strconv.ParseBool("true")
	fmt.Printf("%t\n", boolVal)

	quoted := strconv.Quote("Go rocks!")
	fmt.Printf("%s\n", quoted)

	unquoted, _ := strconv.Unquote(quoted)
	fmt.Printf("%s\n", unquoted)
}
