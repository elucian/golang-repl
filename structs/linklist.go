package structs

import "fmt"

// A single element
type Node struct {
	prev *Node
	next *Node
	value  interface{}
}

// A collection of elements
type List struct {
	head *Node
	tail *Node
}

// List method, used to create elements
func (L *List) Insert(value interface{}) {
	list := &Node{
		next: L.head,
		key:  value,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

// List method
func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.value)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.value)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.value)
		list = list.prev
	}
	fmt.Println()
}

// List method
func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

// main function
func LinkedList() {
  // create empty list
	link := List{}

  // insert elements one by one
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

  //print the list in order
	fmt.Println("\nThe linked list (link) \n")
	fmt.Printf("Head: %v\n", link.head.value)
	fmt.Printf("Tail: %v\n", link.tail.value)
	link.Display()
  fmt.Println()
  fmt.Println("\nReversed list:\n")
	fmt.Printf("head: %v\n", link.head.value)
	fmt.Printf("tail: %v\n", link.tail.value)
	link.Reverse()
  
	fmt.Println("\ndone.\n")
}