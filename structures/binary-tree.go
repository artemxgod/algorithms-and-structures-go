package structures

import (
	"errors"
	"fmt"
)

// Every node of a tree contains value, left and right subtrees
// left subtree should contain value lesser then current tree
// right should contain greater value
type Tree struct {
	val   int
	left  *Tree
	right *Tree
}

func NewTree(rootVal int) *Tree {
	return &Tree{val: rootVal}
}

// Inserting value into a tree node
func (t *Tree) Insert(value int) error {
	if t == nil {
		t = NewTree(value)
	}
	if t.val == value {
		return errors.New("this value already exists")
	}

	if value < t.val {
		if t.left == nil {
			t.left = &Tree{val: value}
			return nil
		}
		return t.left.Insert(value)
	} else {
		if t.right == nil {
			t.right = &Tree{val: value}
			return nil
		}
		return  t.right.Insert(value)
	}
}

// Seeking for tree's maximum value
func (t *Tree) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

// Seeking for tree's minimum value
func (t *Tree) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}

func (t *Tree) PrintTree() {
	t.printTreeCorrectly(t.FindMax())
}

func (t *Tree) printTreeCorrectly(maxval int) {
	if t == nil { return }

	t.left.printTreeCorrectly(maxval)
	if t.val == maxval {
		fmt.Println(t.val)
	} else {
		fmt.Print(t.val, " | ")
	}
	t.right.printTreeCorrectly(maxval)
}



