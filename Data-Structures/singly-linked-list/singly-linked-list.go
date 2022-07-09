package main

import "fmt"

// This is a struct which is a node and contain name and pointer to the next element
type ele struct {
	name string
	next *ele
}

// This is struct which contain lenght and head parts of singly-linked-list
type singleList struct {
	len  int
	head *ele
}

// To initialize the singly-linked-list
func initList() *singleList {
	return &singleList{}
}

// To add element in front of singly-linked-list, AddFront is *singleList type
func (s *singleList) AddFront(name string) {
	// Take element name of string type
	ele := &ele{
		name: name,
	}
	// Check if head pointer is nil, which mean list is empty
	if s.head == nil {
		// If head of list is empty, then set head pointer to element
		s.head = ele
	} else {
		// if list element is not empty then simply update the position of head to next pointer and head to element
		ele.next = s.head
		s.head = ele
	}
	// Increase length of list
	s.len++
	return
}

// This function add element in the back of the list
func (s *singleList) AddBack(name string) {
	// Take element name of the string type
	ele := &ele{
		name: name,
	}
	// Check if list is empty, by checking if head is nill
	if s.head == nil {
		// If the head is empty we will simply update the pointer of the head
		s.head = ele
	} else {
		// Traverse to the last element
		current := s.head
		for current.next != nil {
			current = current.next
		}
		// Set the last element head to the element
		current.next = ele
	}
	// Increase the length of list
	s.len++
	return
}

// This function remove element from the Front i.e. head
func (s *singleList) RemoveFront() error {
	// Check if list is empty
	if s.head == nil {
		return fmt.Errorf("List is empty")
	}
	// If is list is not empty, change head pointer to the next element
	s.head = s.head.next
	// Decrease len of list by 1
	s.len--
	return nil
}

// Remove element from the back
func (s *singleList) RemoveBack() error {
	// Check if the list is empty, if it does throw error
	if s.head == nil {
		return fmt.Errorf("removeBack: List is empty")
	}

	// Traverse through the list and check pointer of two element, 1. if current is last then prev is 2nd last element
	var prev *ele
	current := s.head
	for current.next != nil {
		prev = current
		current = current.next
	}
	// If previous element is not nil, set it to nil
	if prev != nil {
		prev.next = nil
	} else {
		// Head is set to nil
		s.head = nil
	}
	// Decrease length by 1
	s.len--
	return nil
}

// This function gives element from the front
func (s *singleList) Front() (string, error) {
	// If list is empty then head would be nil.
	if s.head == nil {
		return "", fmt.Errorf("Single List is empty")
	}
	// else return the head element
	return s.head.name, nil
}

// This function gives size of list
func (s *singleList) Size() int {
	return s.len
}

// This function traverse through the list and print each element
func (s *singleList) Traverse() error {
	// If head is empty through the error
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	// Set current element to head and print each element until it is equal to empty
	current := s.head
	for current != nil {
		fmt.Println(current.name)
		current = current.next
	}
	return nil
}

// GetAt returns node at given position from linked list
func (l *singleList) GetAt(pos int) *ele {
	// Take current position of header
	ptr := l.head

	// If position is less than 0, through error that it is not a valid position
	if pos < 0 {
		return ptr
	}

	// If position is greater than the list size, through error that position exceed the list size.
	if pos > (l.len - 1) {
		return nil
	}

	// If position is valid, traverse through the list by updating the head pointer.
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// InsertAt inserts new node at given position
func (l *singleList) InsertAt(pos int, name string) {
	// create a new node
	// newNode := Node{}
	ele := &ele{
		name: name,
	}
	// newNode.value = value
	// If position is less than the 0, than through error that it is not a valid position.
	if pos < 0 {
		return
	}

	// If position is 0, insert element at front and and increase list size by 1.
	if pos == 0 {
		l.head = ele
		l.len++
		return
	}
	// If position is greater than the list size throw the error that it is not a valid position.
	if pos > l.len {
		return
	}

	// Traverse through the list and get the element at the givin position.
	n := l.GetAt(pos)
	// Set the element position to the next pointer.
	ele.next = n
	// Get the element to the previous position.
	prevNode := l.GetAt(pos - 1)
	// Set the privious position node pointer the current element
	prevNode.next = ele
	// increase lenght by 1
	l.len++
}

func main() {
	singleList := initList()
	fmt.Printf("AddFront: A\n")
	singleList.AddFront("A")
	fmt.Printf("AddFront: B\n")
	singleList.AddFront("B")
	fmt.Printf("AddBack: C\n")
	singleList.AddBack("C")

	fmt.Printf("Size: %d\n", singleList.Size())

	err := singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveFront\n")
	err = singleList.RemoveFront()
	if err != nil {
		fmt.Printf("RemoveFront Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Size: %d\n", singleList.Size())
}
