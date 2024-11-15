package main

import "fmt"

type Node struct {
	item string
	next *Node
}

type linkedList struct {
	head *Node
	size int
}

func (p *linkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currentNode := p.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	p.size += 1
	return nil
}

func (p *linkedList) removeNode(index int) error {
	if p.head == nil {
		fmt.Println("Linked list is empty.")
	}
	currentNode := p.head
	if index == 0 {
		p.head = currentNode.next
	} else {
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
		nextNode := currentNode.next
		currentNode.next = nextNode.next
	}
	p.size--
	return nil
}

func (p *linkedList) addAtPos(index int, name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
		return nil
	}
	currentNode := p.head
	if index == 0 {
		p.head = newNode
	} else {
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
		nextNode := currentNode.next
		currentNode.next = newNode
		newNode.next = nextNode

	}
	p.size += 1
	return nil
}

func (p *linkedList) get(index int) (Node, error) {
	currentNode := p.head
	if p.head == nil {
		fmt.Println("Linked list is empty.")
		return *currentNode, nil
	}
	if index > 0 {
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
	}
	return *currentNode.next, nil
}

func (p *linkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Linked list is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func main() {
	list := linkedList{
		head: nil,
		size: 0,
	}
	fmt.Println("Activity 1")
	list.addNode("hello")
	list.addNode("world")
	list.addNode("hello2")
	list.addNode("world2")
	list.printAllNodes()
	fmt.Println("-----")
	fmt.Println("Activity 2: Remove")
	list.removeNode(1)
	list.printAllNodes()
	fmt.Println("-----")
	fmt.Println("Activity 2: Add at Position")
	list.addAtPos(2, "inserted")
	list.printAllNodes()
	fmt.Println("-----")
	fmt.Println("Activity 2: Get at Position")
	item, _ := list.get(2)
	fmt.Printf("%v\n", item.item)
}
