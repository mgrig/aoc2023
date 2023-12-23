package day23

type Node struct {
	id   int
	size int
	next []int
}

func NewNode(id int) *Node {
	return &Node{
		id:   id,
		size: 0,
		next: make([]int, 0),
	}
}

func (n *Node) AddNext(nextId int) {
	n.next = append(n.next, nextId) // TODO not checking if it already exists
}
