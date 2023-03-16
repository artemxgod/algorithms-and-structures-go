package structures

import (
	"fmt"
	"strings"
)

type VNode struct {
	Value int
}

type Stack struct {
	nodes []*VNode
	count int
}

func (n *VNode) String() string {
	return fmt.Sprint(n.Value)
}

// Stack is a basic LIFO stack that resizes as needed.
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(n *VNode) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *VNode {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func PrintStack(s *Stack) {
	var out string
	for _, elem := range s.nodes {
		out += fmt.Sprint(elem, "->")
	}
	out = strings.TrimSuffix(out, "->")
	fmt.Println(out)
}

func TestStack() {
	s := NewStack()
	s.Push(&VNode{3})
	s.Push(&VNode{5})
	s.Push(&VNode{7})
	s.Push(&VNode{9})
	PrintStack(s)
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
}
