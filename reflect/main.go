package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string
	Age   int
	Score float64
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
}

func main() {
	person := Person{"Aryn", 22, 90.5}

	personType := reflect.TypeOf(person)
	fmt.Printf("%v\n", personType)

	ageField := reflect.ValueOf(person).FieldByName("Age")
	fmt.Printf("%d\n", ageField.Int())

	reflect.ValueOf(&person).Elem().FieldByName("Name").SetString("Bima")
	fmt.Printf("%s\n", person.Name)

	newPerson := reflect.New(personType).Elem().Interface().(Person)
	fmt.Printf("%v\n", newPerson)

	methodResult := reflect.ValueOf(person).MethodByName("GetInfo").Call([]reflect.Value{})
	fmt.Printf("%v\n", methodResult[0].Interface())
}
