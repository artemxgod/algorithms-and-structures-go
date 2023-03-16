package structures

import "fmt"

// In computer science, a linked list is a linear collection of data elements,
// in which linear order is not given by their physical placement in memory.
// Each pointing to the next node by means of a pointer.
// It is a data structure consisting of a group of nodes which together represent a sequence.
// Here is source code of the Go Program to Implement Single Unsorted Linked List

type Node struct {
	prev *Node
	next *Node
	key interface{}
}

type List struct {
	Head *Node
	Tail *Node
}

func (L *List) Insert(p_key interface{}) {
	// creating a new node, it goes to the beggining, so the next element is the head
	list := &Node{
		next: L.Head,
		key: p_key,
	}

	// connecting with the prev elemenet
	if L.Head != nil {
		L.Head.prev = list
	}
	// pointing head element on new added element
	L.Head = list

	// seeking for the tail
	l := L.Head
	for l.next != nil {
		l = l.next
	}

	// determine the tail
	L.Tail = l
}

func (l *List) Display() {
	list := l.Head
	for list != nil {
		fmt.Printf("%+v->", list.key)
		list = list.next
	}
	fmt.Printf("%s\n", "nil")
}

func PrintList(list *Node) {
	for list != nil {
		fmt.Printf("%v->", list.key)
		list = list.next
	}
	fmt.Printf("%s\n", "nil")
}

func PrintListBackwords(list *Node) {
	for list != nil {
		fmt.Printf("%v<-", list.key)
		list = list.prev
	}
	fmt.Printf("%s\n", "nil")
}

func (l *List) Reverse() {
	cur := l.Head
	var prev *Node
	l.Tail = l.Head

	for cur != nil {
		next := cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}
	l.Head = prev
	
}

func (l *List) Push(key interface{}) {
	list := &Node{
		prev: l.Tail,
		key: key,
	}

	if l.Tail != nil {
		l.Tail.next = list
	}
	l.Tail = list
	hd := l.Tail
	for hd.prev != nil {
		hd = hd.prev
	}

	l.Head = hd
}

func (l *List) Pop() {
	del := l.Tail
	l.Tail = l.Tail.prev
	del.prev = nil
	l.Tail.next = nil
}

func (l *List) Sort() {
	cur := l.Head.next

	for cur != nil {
		next := cur.next
		for next != nil && next.prev != nil {
			if next.prev.key.(int) > next.key.(int) {
				next.prev.key, next.key = next.key, next.prev.key
			}
			next = next.prev
		}
		cur = cur.next
	}
}

func TestLinkedList() {
	list := List{}

	fmt.Println("=====INSERTED=====")
	list.Insert(4)
	list.Insert(5)
	list.Insert(123)
	list.Insert(8)
	PrintList(list.Head)

	fmt.Println("=====INSERTED BACKWORDS=====")
	PrintListBackwords(list.Tail)

	fmt.Println("=====REVERSED=====")
	list.Reverse()
	list.Display()

	fmt.Println("=====REVERSED BACKWORDS=====")
	PrintListBackwords(list.Tail)

	fmt.Println("=====PUSHED=====")
	list.Push("Amen")
	PrintList(list.Head)

	fmt.Println("=====PUSHED BACKWORDS=====")
	PrintListBackwords(list.Tail)

	fmt.Println("=====POPPED=====")
	list.Pop()
	list.Display()

	fmt.Println("=====POPPED BACKWORDS=====")
	PrintListBackwords(list.Tail)

	fmt.Println("=====SORTED LIST=====")
	list.Sort()
	list.Display()
}
