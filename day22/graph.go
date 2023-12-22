package day22

type Node struct {
	id          int
	supports    map[int]bool
	supportedBy map[int]bool
}

func NewNode(id int) *Node {
	return &Node{
		id:          id,
		supports:    make(map[int]bool, 0),
		supportedBy: make(map[int]bool, 0),
	}
}

func (n *Node) Supports(other int) {
	n.supports[other] = true
}

func (n *Node) SupportedBy(other int) {
	n.supportedBy[other] = true
}

func (n *Node) NextNodes() []int {
	ret := make([]int, len(n.supports))
	i := 0
	for next := range n.supports {
		ret[i] = next
		i++
	}
	return ret
}

func (n *Node) PrevNodes() []int {
	ret := make([]int, len(n.supportedBy))
	i := 0
	for next := range n.supportedBy {
		ret[i] = next
		i++
	}
	return ret
}
