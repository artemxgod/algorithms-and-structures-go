package structures

import (
	"fmt"
	"strings"
)

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
    nodes []*VNode
    size  int
    head  int
    tail  int
    count int
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
    return &Queue{
        nodes: make([]*VNode, size),
        size:  size,
    }
}

// Push adds a node to the queue.
func (q *Queue) Push(n *VNode) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*VNode, len(q.nodes) + q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes) - q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func (q *Queue) Pop() *VNode {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func PrintQueue(q *Queue) {
	var out string
	for _, elem := range q.nodes {
		out += fmt.Sprint(elem, "->")
	}
	out = strings.TrimSuffix(out, "->")
	fmt.Println(out)
}

func TestQueue() {
	q := NewQueue(1)
    q.Push(&VNode{2})
    q.Push(&VNode{4})
    q.Push(&VNode{6})
    q.Push(&VNode{8})
	PrintQueue(q)
    fmt.Println(q.Pop(), q.Pop(), q.Pop(), q.Pop())
}