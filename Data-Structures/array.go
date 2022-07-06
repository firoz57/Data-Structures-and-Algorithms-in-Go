package main

import "fmt"

func main() {
	// Declaring the variable of size 5 with elements 5,6,3,2
	array := [5]int{5, 6, 3, 2}

	// Printing memory location of each elements.
	for i, element := range array {
		fmt.Print("Element ", element, " having ", &array[i], " Memory location \n")
	}

	// Append into array
	appended_array := make([]int, 10)
	appended_array = append(array[:], 9)
	fmt.Print("This is appended_array", appended_array, "\n")
}
