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
