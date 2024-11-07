package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type queue struct {
	front *Node
	back  *Node
	size  int
}

type stack struct {
	top  *Node
	size int
}

func (p *stack) push(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top
		p.top = newNode
	}
	p.size++
	return nil
}

func (p *stack) pop() (string, error) {
	var item string

	if p.top == nil {
		return "", errors.New("Empty Stack!")
	}

	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

func (p *stack) printAllNodes() error {
	currentNode := p.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	return nil
}

func (p *queue) enqueue(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.front == nil {
		p.front = newNode

	} else {
		p.back.next = newNode

	}
	p.back = newNode
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

func (p *queue) isEmpty() bool {
	return p.size == 0
}

func checkPalindrome(input string) {
	paliQueue := queue{
		front: nil,
		back:  nil,
		size:  0,
	}
	paliStack := stack{
		top:  nil,
		size: 0,
	}

	for _, char := range input {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			paliQueue.enqueue(string(char))
			paliStack.push(string(char))
		}
	}

	for paliStack.size > 0 {

		stackValue, err := paliStack.pop()
		if err != nil {
			fmt.Println("error popping from stack")
		}

		queueValue, err := paliQueue.dequeue()
		if err != nil {
			fmt.Println("error dequeueing from queue")
		}
		if stackValue != queueValue {
			fmt.Println("This word is not a palindrome.")
			return
		}
	}
	fmt.Println("This word is a palindrome.")
}

func main() {
	queueOne := queue{
		front: nil,
		back:  nil,
		size:  0,
	}

	fmt.Println("Activity 1")
	queueOne.enqueue("hello")
	queueOne.enqueue("world")
	queueOne.enqueue("hello2")
	queueOne.enqueue("world2")
	queueOne.printAllNodes()
	fmt.Println("--------")

	fmt.Println("Activity 2")
	checkPalindrome("top spot")
	fmt.Println("--------")

}
