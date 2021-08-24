package datastructure

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func createNode(value int) *Node {
	newNode := &Node{value, nil, nil}
	return newNode
}

func (list *LinkedList) pushHead(value int) {
	newNode := createNode(value)
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

func (list *LinkedList) pushTail(value int) {
	newNode := createNode(value)
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

func (list *LinkedList) pushMid(value int) {
	newNode := createNode(value)
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else if value < list.head.value {
		list.pushHead(value)
	} else if value > list.tail.value {
		list.pushTail(value)
	} else {
		currentNode := list.head
		for value > currentNode.value {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode.prev
		newNode.next = currentNode
		currentNode.prev.next = newNode
		currentNode.prev = newNode
	}
}

func (list *LinkedList) popHead() {
	if list.head == nil {
		return
	} else if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		replacingHead := list.head.next
		list.head.next = nil
		list.head.prev = nil
		list.head = replacingHead
	}
}

func (list *LinkedList) popTail() {
	if list.head == nil {
		return
	} else if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		replacingTail := list.tail.prev
		replacingTail.next = nil
		list.tail.prev = nil
		list.tail = replacingTail
	}
}

func (list *LinkedList) popMid(value int) {
	if list.head == nil {
		return
	} else if list.head.value == value {
		list.popHead()
	} else if list.tail.value == value {
		list.popTail()
	} else {
		toBeRemoved := list.head.next

		for value != toBeRemoved.value {
			toBeRemoved = toBeRemoved.next
			if toBeRemoved == list.tail {
				fmt.Println("Not Found")
				return
			}
		}

		toBeRemoved.prev.next = toBeRemoved.next
		toBeRemoved.next.prev = toBeRemoved.prev
		toBeRemoved.next = nil
		toBeRemoved.prev = nil
		toBeRemoved = nil
	}
}

func (list *LinkedList) popAll() {
	currentNode := list.head
	for currentNode != nil {
		list.head = list.head.next
		currentNode = nil
		currentNode = list.head
	}
}

func (list *LinkedList) printAll() {

	if list.head == nil {
		fmt.Println("Linked List is Empty !")
		return
	}

	currentNode := list.head
	for currentNode != nil {
		fmt.Print(currentNode.value)
		if currentNode.next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Println()
		}
		currentNode = currentNode.next
	}
}

func (list LinkedList) simulateOperation(items []int, totalItem int) {
	fmt.Println("INSERTION")
	for i := 0; i < totalItem; i++ {
		list.pushMid(items[i])
		fmt.Println("Inserted :", items[i])
		fmt.Println("Current Linked List : ")
		list.printAll()
		fmt.Println("Press enter to continue the next operation ...")
		fmt.Scanln()
	}
	fmt.Println("DELETION")
	for i := 0; i < totalItem; i++ {
		list.popMid(items[i])
		fmt.Println("Removed :", items[i])
		fmt.Println("Current Linked List : ")
		list.printAll()
		fmt.Println("Press enter to continue the next operation ...")
		fmt.Scanln()
	}
	fmt.Println("FINISHED")
}

func RunLinkedList() {
	linkedList := &LinkedList{}
	items := []int{13, 12, 16, 15, 8, 14, 11, 18}
	totalItem := len(items)
	linkedList.simulateOperation(items, totalItem)
	linkedList.popAll()
}
