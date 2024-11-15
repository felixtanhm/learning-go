package main

import (
	"fmt"
)

type BinaryNode struct {
	item  string      // to store the data item
	left  *BinaryNode // pointer to point to left node
	right *BinaryNode // pointer to point to right node
}

type BST struct {
	root *BinaryNode
}

func (bst *BST) insertNode(t **BinaryNode, item string) error {

	if *t == nil {
		newNode := &BinaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}

	if item < (*t).item {
		bst.insertNode(&((*t).left), item)
	} else {
		bst.insertNode(&((*t).right), item)
	}

	return nil
}

func (bst *BST) insert(item string) {
	bst.insertNode(&bst.root, item)
}

func (bst *BST) preOrder() {
	bst.traverse(bst.root, "preOrder", bst.visitNode)
}

func (bst *BST) inOrder() {
	bst.traverse(bst.root, "inOrder", bst.visitNode)
}

func (bst *BST) postOrder() {
	bst.traverse(bst.root, "postOrder", bst.visitNode)
}

func (bst *BST) traverse(t *BinaryNode, order string, visit func(*BinaryNode)) {
	if t == nil {
		return
	}

	if order == "preOrder" {
		visit(t)
	}
	bst.traverse(t.left, order, visit)
	if order == "inOrder" {
		visit(t)
	}
	bst.traverse(t.right, order, visit)
	if order == "postOrder" {
		visit(t)
	}
}

func (bst *BST) visitNode(t *BinaryNode) {
	fmt.Println(t.item)
}

func (bst *BST) getNode(t *BinaryNode, name string) *BinaryNode {
	if t == nil {
		return nil
	}
	if t.item == name {
		return t
	}

	left := bst.getNode(t.left, name)
	if left != nil {
		return left
	}

	return bst.getNode(t.right, name)
}

func (bst *BST) get(name string) *BinaryNode {
	return bst.getNode(bst.root, name)
}

func (bst *BST) countNodes(t *BinaryNode) int {
	if t == nil {
		return 0
	}

	return 1 + bst.countNodes(t.left) + bst.countNodes(t.right)
}

func (bst *BST) count() int {
	return bst.countNodes(bst.root)
}

func (bst *BST) removeNode(t *BinaryNode, name string) *BinaryNode {
	if t == nil {
		return nil
	}

	if name < t.item {
		t.left = bst.removeNode(t.left, name)
	} else if name > t.item {
		t.right = bst.removeNode(t.right, name)
	} else {
		if t.left == nil {
			return t.right
		} else if t.right == nil {
			return t.left
		} else {
			successor := t.right
			for successor.left != nil {
				successor = successor.left
			}
			t.item = successor.item
			t.right = bst.removeNode(t.right, t.item)
		}
	}
	return t
}

func (bst *BST) remove(name string) {
	bst.removeNode(bst.root, name)
}

func main() {
	tree := BST{
		root: nil,
	}

	// fmt.Println("Activity 1")
	tree.insert("5")
	tree.insert("2")
	tree.insert("4")
	tree.insert("3")
	tree.insert("6")
	tree.insert("1")
	tree.insert("8")
	tree.insert("9")
	tree.insert("7")
	tree.inOrder()
	fmt.Println("----------")

	fmt.Println("Activity 2")
	fmt.Println(tree.get("6").item)
	fmt.Println("----------")

	fmt.Println("Activity 3")
	fmt.Println("Pre Order Traversal:")
	tree.preOrder()
	fmt.Println("Post Order Traversal:")
	tree.postOrder()
	fmt.Println("----------")

	fmt.Println("Activity 4")
	fmt.Println(tree.count())
	fmt.Println("----------")

	fmt.Println("Activity 5")
	tree.remove("6")
	fmt.Println("----------")
}
