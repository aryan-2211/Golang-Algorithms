package LinkedList

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{val: data}

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

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("the list is empty")
		return
	}

	for current != nil {
		fmt.Printf("%d ", current.val)
		current = current.next
	}
	fmt.Println()

}
