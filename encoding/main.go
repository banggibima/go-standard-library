package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func main() {
	message := "Hello, Go!"
	encoded := base64.StdEncoding.EncodeToString([]byte(message))
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Printf("%s\n", encoded)
	fmt.Printf("%s\n", decoded)

	person := Person{"Aryn", 22, 92.5}
	jsonData, _ := json.Marshal(person)
	fmt.Printf("%s\n", jsonData)

	var newPerson Person
	json.Unmarshal(jsonData, &newPerson)
	fmt.Printf("%+v\n", newPerson)

	csvData := "John,25,90.5,Jane,30,85.0"
	reader := csv.NewReader(strings.NewReader(csvData))
	records, _ := reader.ReadAll()
	fmt.Printf("%+v\n", records)

	var csvBuilder strings.Builder
	csvWriter := csv.NewWriter(&csvBuilder)
	csvWriter.Write([]string{"Alice", "28", "88.2"})
	csvWriter.Write([]string{"Bob", "35", "76.5"})
	csvWriter.Flush()
	csvContent := csvBuilder.String()
	fmt.Printf("%s", csvContent)
}
