package day23

type Nodes map[int]Node

func NewNodes() *Nodes {
	var ret Nodes = map[int]Node{}
	return &ret
}

func (n *Nodes) GetOrCreate(nodeId int) Node {
	node, exists := (*n)[nodeId]
	if !exists {
		node = *NewNode(nodeId)
		(*n)[nodeId] = node
	}
	return node
}

func (n *Nodes) Set(node Node) {
	(*n)[node.id] = node
}

func (n *Nodes) GetEndNode() Node {
	for _, node := range *n {
		if len(node.next) == 0 {
			return node
		}
	}
	panic("end node not found")
}
