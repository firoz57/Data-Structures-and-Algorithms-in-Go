package main

import (
	"container/list"
	"fmt"
)

// List is essentially a doubly linked list in golang. This not a standard implimentation but a library of golang.
func main() {
	//Initializing of list
	l := list.New()
	fmt.Print("Printing list: ", l, "\n")

	// Fetching last element of list
	fmt.Print("The Last element of list is: ", l.Back(), "\n")

	// Fetching first element of the list
	fmt.Print("The Front element of the list is: ", l.Front(), "\n")

	// Inserting element in the first position
	a := l.PushFront(77)
	// Inserting element in the last position
	b := l.PushBack(33)
	c := l.PushBack(44)
	d := l.PushBack(55)
	e := l.PushBack(66)
	f := l.PushBack(77)
	g := l.PushBack(88)

	fmt.Print("few reference  are", a, b, c, d, e, f, g)
	// Printing first and last values of the list
	fmt.Printf("New element is added in front %d and back %d : \n", l.Front().Value, l.Back().Value)
	fmt.Print("Printing list: ", l, "\n")

	// Removing the element
	l.Remove(l.Front())
	fmt.Printf("This will remove the front element of list which is 77 and will only show 33 -> %d with length %d: ", l.Front().Value, l.Len())

	// Lookup for at the index in list
	l.Remove(f)
	fmt.Print("\nThis will remove the element 77 from the list and reduce length to: ", l.Len())
}
