package models

import (
	"strings"
	"sync"
)

type Venue struct {
	Uuid        string
	Name        string
	Type        string
	Capacity    int
	BookingList LinkedList
}

type BinaryNode struct {
	venue Venue       // to store the data item
	left  *BinaryNode // pointer to point to left node
	right *BinaryNode // pointer to point to right node
}

type BST struct {
	root  *BinaryNode
	mutex sync.Mutex
}

func CreateBST() *BST {
	return &BST{root: nil}
}

func (bst *BST) Insert(venue Venue) {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	bst.insertNode(&bst.root, venue)
}

func (bst *BST) GetAll(order string) []*Venue {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	limit := int(bst.countNodes(bst.root))
	results := []*Venue{}
	startIndex := 0
	bst.traverse(bst.root, order, &results, &startIndex, limit)

	return results
}

func (bst *BST) GetList(order string, page int, limit int) []*Venue {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	results := []*Venue{}
	startIndex := (page - 1) * limit
	bst.traverse(bst.root, order, &results, &startIndex, limit)

	return results
}

func (bst *BST) GetOne(name string) *Venue {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	result := bst.getNode(bst.root, name)
	if result != nil {
		return &result.venue
	} else {
		return nil
	}
}

func (bst *BST) Remove(name string) {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	bst.removeNode(bst.root, name)
}

func (bst *BST) Count() float64 {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	return bst.countNodes(bst.root)
}

func (bst *BST) insertNode(t **BinaryNode, venue Venue) error {
	if *t == nil {
		newNode := &BinaryNode{
			venue: venue,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}

	if venue.Name < (*t).venue.Name {
		bst.insertNode(&((*t).left), venue)
	} else {
		bst.insertNode(&((*t).right), venue)
	}

	return nil
}

func (bst *BST) traverse(t *BinaryNode, order string, results *[]*Venue, startIndex *int, limit int) {
	if t == nil || len(*results) >= limit {
		return
	}

	if order == "preOrder" {
		bst.visitNode(t, results, startIndex, limit)
	}
	bst.traverse(t.left, order, results, startIndex, limit)
	if order == "inOrder" {
		bst.visitNode(t, results, startIndex, limit)
	}
	bst.traverse(t.right, order, results, startIndex, limit)
	if order == "postOrder" {
		bst.visitNode(t, results, startIndex, limit)
	}
}

func (bst *BST) visitNode(t *BinaryNode, results *[]*Venue, startIndex *int, limit int) {
	if *startIndex > 0 {
		*startIndex--
	} else if len(*results) < limit {
		*results = append(*results, &t.venue)
	}
}

func (bst *BST) getNode(t *BinaryNode, name string) *BinaryNode {
	if t == nil {
		return nil
	}

	if strings.EqualFold(t.venue.Name, name) {
		return t
	}

	left := bst.getNode(t.left, name)
	if left != nil {
		return left
	}

	return bst.getNode(t.right, name)
}

func (bst *BST) countNodes(t *BinaryNode) float64 {
	if t == nil {
		return 0
	}

	return 1 + bst.countNodes(t.left) + bst.countNodes(t.right)
}

func (bst *BST) removeNode(t *BinaryNode, name string) *BinaryNode {
	if t == nil {
		return nil
	}

	if name < t.venue.Name {
		t.left = bst.removeNode(t.left, name)
	} else if name > t.venue.Name {
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
			t.venue = successor.venue
			t.right = bst.removeNode(t.right, t.venue.Name)
		}
	}
	return t
}
