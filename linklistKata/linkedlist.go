package linklistKata

import "fmt"

// Node represents an element in the linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Insert adds a new node to the end of the list
func (list *LinkedList) Insert(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Delete removes a node by value
func (list *LinkedList) Delete(value int) {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}
	if list.head.value == value {
		list.head = list.head.next
		return
	}
	current := list.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next == nil {
		fmt.Println("Value not found in the list")
	} else {
		current.next = current.next.next
	}
}

// Display prints all the nodes in the list
func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
