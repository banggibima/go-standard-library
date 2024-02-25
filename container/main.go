package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func contains(l *list.List, value interface{}) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true
		}
	}
	return false
}

func main() {
	myList := list.New()

	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%v\n", element.Value)
	}

	removeElement := myList.Front()
	myList.Remove(removeElement)

	containsValue := contains(myList, 2)
	fmt.Printf("%t\n", containsValue)

	listLength := myList.Len()
	fmt.Printf("%d\n", listLength)

	myRing := ring.New(3)

	for i := 1; i <= 3; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Do(func(value interface{}) {
		fmt.Printf("%v\n", value)
	})

	myRing = myRing.Move(2)

	currentValue := myRing.Value
	fmt.Printf("%v\n", currentValue)

	anotherRing := ring.New(2)
	myRing.Link(anotherRing)

	unlinkedRing := myRing.Unlink(2)

	unlinkedRing.Do(func(value interface{}) {
		fmt.Printf("%v\n", value)
	})
}
