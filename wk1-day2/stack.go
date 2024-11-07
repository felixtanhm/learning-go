package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
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

func checkBalance(input string) (bool, error) {
	checkStack := stack{
		top:  nil,
		size: 0,
	}

	for _, char := range input {
		if char == '(' || char == '[' || char == '{' {
			checkStack.push(string(char))
		} else {
			poppedChar, err := checkStack.pop()
			if err != nil {
				fmt.Println("error popping from check balance stack")
			}
			switch poppedChar {
			case "(":
				if string(char) != ")" {
					return false, nil
				}
			case "{":
				if string(char) != "}" {
					return false, nil
				}
			case "[":
				if string(char) != "]" {
					return false, nil
				}
			}
		}
	}
	if checkStack.size > 0 {
		return false, nil
	}

	return true, nil
}

func main() {
	stackOne := stack{
		top:  nil,
		size: 0,
	}

	fmt.Println("Activity 1")
	stackOne.push("hello")
	stackOne.push("world")
	stackOne.push("hello2")
	stackOne.push("world2")
	stackOne.printAllNodes()
	fmt.Println("------")

	fmt.Println("Activity 2")
	tempStack := stack{
		top:  nil,
		size: 0,
	}
	for stackOne.size > 0 {
		element, err := stackOne.pop()
		if err != nil {
			fmt.Println("error popping from stack")
		}
		err = tempStack.push(element)
		if err != nil {
			fmt.Println("error pushing to temp stack")
		}
	}
	tempStack.printAllNodes()
	fmt.Println("------")

	fmt.Println("Activity 3")
	isBalanced, err := checkBalance("({[])")
	if err != nil {
		fmt.Println("%w", err)
	}
	fmt.Println(isBalanced)
}
