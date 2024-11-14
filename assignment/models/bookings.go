package models

import (
	"fmt"
)

type Booking struct {
	Uuid      string
	VenueID   string
	UserEmail string
}

type LinkedListNode struct {
	booking Booking
	next    *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
	size int
}

func (p *LinkedList) addNode(booking Booking) error {
	newNode := &LinkedListNode{
		booking: booking,
		next:    nil,
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

func (p *LinkedList) removeNode(index int) error {
	if p.head == nil {
		fmt.Println("Bookings for this venue is empty.")
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

func (p *LinkedList) addAtPos(index int, booking Booking) error {
	newNode := &LinkedListNode{
		booking: booking,
		next:    nil,
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

func (p *LinkedList) get(index int) (LinkedListNode, error) {
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

func (p *LinkedList) printAllNodes() error {
	currentNode := p.head
	if currentNode == nil {
		fmt.Println("Bookings for this venue is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.booking)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.booking)
	}

	return nil
}
