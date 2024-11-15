package models

import (
	"fmt"
	"sync"
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
	head  *LinkedListNode
	tail  *LinkedListNode
	size  int
	mutex sync.Mutex
}

func CreateLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil, size: 0}
}

func (p *LinkedList) Insert(booking Booking) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	newNode := &LinkedListNode{
		booking: booking,
		next:    nil,
	}
	if p.head == nil {
		p.head = newNode
		p.tail = newNode
	} else {
		p.tail.next = newNode
		p.tail = newNode
	}
	p.size++
	return nil
}

func (p *LinkedList) Remove(uuid string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.head == nil {
		return fmt.Errorf("bookings for this venue is empty")
	}
	var prevNode *LinkedListNode
	currentNode := p.head

	for currentNode != nil {
		if uuid == currentNode.booking.Uuid {
			if prevNode == nil {
				p.head = currentNode.next
				if p.head == nil {
					p.tail = nil
				}
			} else {
				prevNode.next = currentNode.next
				if currentNode == p.tail {
					p.tail = prevNode
				}
			}
			p.size--
			return nil
		}
		prevNode = currentNode
		currentNode = currentNode.next
	}
	p.size--
	return fmt.Errorf("booking with UUID %v not found", uuid)
}

func (p *LinkedList) GetAll() []*Booking {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	var results []*Booking
	currentNode := p.head
	if currentNode == nil {
		return results
	}
	results = append(results, &currentNode.booking)
	for currentNode.next != nil {
		currentNode = currentNode.next
		results = append(results, &currentNode.booking)
	}
	return results
}
