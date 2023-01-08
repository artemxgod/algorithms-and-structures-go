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
		return t.right.Insert(value)
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

// Seeking for tree's minimum value node
func (t *Tree) GetMin() *Tree {
	if t.left == nil {
		return t
	}
	return t.left.GetMin()
}

func (t *Tree) PrintTree() {
	if t == nil {
		return
	}
	t.printTreeCorrectly(t.FindMax())
}
func (t *Tree) printTreeCorrectly(maxval int) {
	if t == nil {
		return
	}

	t.left.printTreeCorrectly(maxval)
	if t.val == maxval {
		fmt.Println(t.val)
	} else {
		fmt.Print(t.val, " | ")
	}
	t.right.printTreeCorrectly(maxval)
}

func (t *Tree) PrintPreorder() {
	if t == nil {
		return
	}

	fmt.Print(t.val, " | ")
	t.left.PrintPreorder()
	t.right.PrintPreorder()
}

// seeking for node and it's parent
func (t *Tree) FindNodeAndParent(value int) (node, parent *Tree) {
	if t.val == value {
		return t, nil
	}
	return t.findNodeAndParentReq(value)
}
func (t *Tree) findNodeAndParentReq(value int) (node, parent *Tree) {
	if value < t.val {
		if t.left == nil {
			return nil, nil
		}
		if t.left.val == value {
			return t.left, t
		} else {
			return t.left.findNodeAndParentReq(value)
		}
	} else {
		if t.right == nil {
			return nil, nil
		}
		if t.right.val == value {
			return t.right, t
		} else {
			return t.right.findNodeAndParentReq(value)
		}
	}

}

func (t *Tree) DeleteNode(value int) (*Tree, error) {
	// searching for node, and check if it was found
	node, parent := t.FindNodeAndParent(value)
	if node == nil && parent == nil {
		return t, errors.New("node not found")
	}

	if node.left == nil && node.right == nil {
		return t.deleteNodeNoSuccessors(parent, node, value), nil
	} 
	if node.left != nil && node.right == nil {
		return t.deleteNodeOneSuccessors(parent, node, node.left, value), nil
	}
	if node.left == nil && node.right != nil {
		return t.deleteNodeOneSuccessors(parent, node, node.right, value), nil
	}
	if node.left != nil && node.right != nil {
		successor := node.right.GetMin()
		curVal := successor.val
		t.DeleteNode(curVal)
		node.val = curVal
		return t, nil
	}



	return nil, nil
}

// deleting node with no successors
func (t *Tree) deleteNodeNoSuccessors(parent, node *Tree, value int) *Tree {
	if parent != nil {
		if parent.left != nil && parent.left.val == node.val {
			parent.left = nil
			return t
		} else if parent.right.val == node.val {
			parent.right = nil
			return t
		}
	} else {
		return nil
	}
	return nil
}

// deleting node with one successor
func (t *Tree) deleteNodeOneSuccessors(parent, node, successor *Tree, value int) *Tree {
	if parent != nil {
		if parent.left != nil && parent.left.val == node.val {
			parent.left = successor
			return t
		} else if parent.right.val == node.val {
			parent.right = successor
			return t
		}
	} else {
		return successor
	}
	return nil
}

