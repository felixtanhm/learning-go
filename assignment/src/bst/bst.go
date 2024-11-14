package bst

import (
	"assignment/models"
	"sync"
)

type BinaryNode struct {
	venue models.Venue // to store the data item
	left  *BinaryNode  // pointer to point to left node
	right *BinaryNode  // pointer to point to right node
}

type BST struct {
	root  *BinaryNode
	mutex sync.Mutex
}

func Create() *BST {
	return &BST{root: nil}
}

func (bst *BST) Insert(venue models.Venue) {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	bst.insertNode(&bst.root, venue)
}

func (bst *BST) GetAll(order string) []*models.Venue {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	limit := bst.countNodes(bst.root)
	results := []*models.Venue{}
	startIndex := 0
	bst.traverse(bst.root, order, &results, &startIndex, limit)

	return results
}

func (bst *BST) GetList(order string, page int, limit int) []*models.Venue {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	results := []*models.Venue{}
	startIndex := (page - 1*limit)
	bst.traverse(bst.root, order, &results, &startIndex, limit)

	return results
}

func (bst *BST) GetOne(name string) models.Venue {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	return bst.getNode(bst.root, name).venue
}

func (bst *BST) Remove(name string) {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	bst.removeNode(bst.root, name)
}

func (bst *BST) Count() int {
	bst.mutex.Lock()
	defer bst.mutex.Unlock()

	return bst.countNodes(bst.root)
}

func (bst *BST) insertNode(t **BinaryNode, venue models.Venue) error {
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

func (bst *BST) traverse(t *BinaryNode, order string, results *[]*models.Venue, startIndex *int, limit int) {
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

func (bst *BST) visitNode(t *BinaryNode, results *[]*models.Venue, startIndex *int, limit int) {
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
	if t.venue.Name == name {
		return t
	}

	left := bst.getNode(t.left, name)
	if left != nil {
		return left
	}

	return bst.getNode(t.right, name)
}

func (bst *BST) countNodes(t *BinaryNode) int {
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
