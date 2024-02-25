package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "Buffered Reader Example"
	reader := bufio.NewReader(strings.NewReader(str))
	data, _ := reader.ReadBytes(' ')
	fmt.Printf("%s\n", string(data))

	file, _ := os.Create("./bufio/output.txt")
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("Writing to a file using bufio.Writer.")
	writer.Flush()

	file, _ = os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Printf("%s\n", scanner.Text())
}
