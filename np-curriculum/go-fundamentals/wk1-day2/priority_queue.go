package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item     string
	priority int
	next     *Node
}

type queue struct {
	front *Node
	back  *Node
	size  int
}

func (p *queue) enqueue(name string, priority int) error {
	newNode := &Node{
		item:     name,
		priority: priority,
		next:     nil,
	}
	if p.front == nil {
		p.front = newNode
		p.back = newNode
	} else {
		targetNode := p.getLast(priority)
		if targetNode == nil {
			// There is no Node with a higher || equal priority
			newNode.next = p.front
			p.front = newNode
		} else {
			// Insert after returned target node
			newNode.next = targetNode.next
			targetNode.next = newNode

			if newNode.next == nil {
				// Update back if returned target node is the last node
				p.back = newNode
			}
		}
	}

	p.size++
	return nil
}

func (p *queue) dequeue() (string, error) {
	var item string

	if p.front == nil {
		return "", errors.New("empty queue!")
	}

	item = p.front.item
	if p.size == 1 {
		p.front = nil
		p.back = nil
	} else {
		p.front = p.front.next
	}
	p.size--
	return item, nil
}

func (p *queue) printAllNodes() error {
	currentNode := p.front
	if currentNode == nil {
		fmt.Println("Queue is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *queue) getLast(priority int) *Node {
	currentNode := p.front
	for currentNode.next != nil && currentNode.next.priority <= priority {
		currentNode = currentNode.next
	}
	return currentNode
}

// 1 1 2 2 3

func (p *queue) isEmpty() bool {
	return p.size == 0
}

func main() {
	queueOne := queue{
		front: nil,
		back:  nil,
		size:  0,
	}

	fmt.Println("Activity 3")
	queueOne.enqueue("hello", 1)
	queueOne.enqueue("world", 3)
	queueOne.enqueue("hello2", 7)
	queueOne.enqueue("world2", 5)
	queueOne.enqueue("hello3", 9)
	queueOne.enqueue("world3", 11)
	queueOne.printAllNodes()
	fmt.Println("--------")

}
