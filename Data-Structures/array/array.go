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

	// Insertion into array
	arr1 := []int{2, 6, 8, 1}
	arr1 = append(arr1, 0)
	copy(arr1[3:], arr1[2:])
	arr1[2] = 99
	fmt.Print("This is insertion of 99 in 2nd index: ", arr1, "\n")

	//update of element
	arr1[4] = 77
	fmt.Print("This is update of element 77 at index 4: ", arr1, "\n")

	// deletion of an element from an array
	arr2 := []int{22, 33, 44, 55}
	i := 2

	fmt.Print("This array before the deletion: ", arr2, "\n")

	copy(arr2[i:], arr2[i+1:])
	arr2[len(arr2)-1] = 0     // Erase last element and write 0 in its place
	arr2 = arr2[:len(arr2)-1] // Truncate slice

	fmt.Print("This delete the element from index 2 which is 44: ", arr2, "\n")

	// Lookup at the index in an array
	i = 2
	fmt.Print("The element at position 2 in an array is: ", arr2[i], "\n")
}
