package day08

type Graph struct {
	left  map[string]string
	right map[string]string
}

func NewGraph() *Graph {
	return &Graph{
		left:  make(map[string]string),
		right: make(map[string]string),
	}
}

func (g *Graph) Add(node, left, right string) {
	g.left[node] = left
	g.right[node] = right
}

func (g *Graph) GetLeft(node string) string {
	val, exists := g.left[node]
	if !exists {
		panic("missing mapping")
	}
	return val
}

func (g *Graph) GetRight(node string) string {
	val, exists := g.right[node]
	if !exists {
		panic("missing mapping")
	}
	return val
}
