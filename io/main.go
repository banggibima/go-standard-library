package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, Go!")
	data := make([]byte, 8)
	n, _ := reader.Read(data)
	fmt.Printf("Read from string: %s\n", string(data[:n]))

	file, _ := os.Create("output.txt")
	defer file.Close()
	file.WriteString("Writing to a file.")

	src := strings.NewReader("Source string.")
	dst := new(strings.Builder)
	io.Copy(dst, src)
	fmt.Printf("Copied string: %s\n", dst.String())

	file, _ = os.Open("input.txt")
	defer file.Close()
	fileReader := io.Reader(file)
	fileData := make([]byte, 12)
	n, _ = fileReader.Read(fileData)
	fmt.Printf("Read from file: %s\n", string(fileData[:n]))

	file, _ = os.Create("output.txt")
	defer file.Close()
	fileWriter := io.Writer(file)
	fileWriter.Write([]byte("Writing to a file using io.Writer."))
}
