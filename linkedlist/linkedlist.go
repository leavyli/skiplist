package linkedlist

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Linklist struct {
	head    *Node
	tail    *Node
	current *Node
}

func (link *Linklist) InsertHead(data interface{}) {
	node := &Node{
		data: data,
		next: nil,
	}

	if link.tail == nil && link.head == nil {
		link.tail = node
		link.head = node
		link.current = link.head
		return
	}

	tmp := link.head
	link.head = node
	node.next = tmp
	link.current = link.head
}

func (link *Linklist) HasNext() bool {
	return link.current != nil
}

func (link *Linklist) Next() interface{} {
	data := link.current.data
	link.current = link.current.next

	return data
}

func (link *Linklist) Each() {
	it := link.head
	for it != nil {
		fmt.Println(it.data)
		it = it.next
	}

}
