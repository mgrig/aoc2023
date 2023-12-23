package day23

import "fmt"

type UNode struct {
	id   int
	neis map[int]int // nei -> weight
}

func NewUNode(id int) *UNode {
	return &UNode{
		id:   id,
		neis: make(map[int]int, 0),
	}
}

// ****

type Graph struct {
	nodes map[int]*UNode
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*UNode),
	}
}

func (g *Graph) GetOrCreateUNode(id int) *UNode {
	node, exists := g.nodes[id]
	if !exists {
		node = NewUNode(id)
		g.nodes[id] = node
	}
	return node
}

func (g *Graph) AddEdge(from, to, weight int) {
	fromNode := g.GetOrCreateUNode(from)
	toNode := g.GetOrCreateUNode(to)

	fromNode.neis[to] = weight
	toNode.neis[from] = weight
}

// panic if not found
func (g *Graph) GetWeight(from, to int) int {
	return g.nodes[from].neis[to]
}

func (g *Graph) Simplify() {
	dirty := true
	for dirty {
		dirty = false
		for _, pnode := range g.nodes {
			if len(pnode.neis) == 2 {
				keys := make([]int, 0, len(pnode.neis))
				for k := range pnode.neis {
					keys = append(keys, k)
				}
				nei1 := keys[0]
				nei2 := keys[1]
				w1 := pnode.neis[nei1]
				w2 := pnode.neis[nei2]
				nei1Node := g.nodes[nei1]
				nei2Node := g.nodes[nei2]

				delete(nei1Node.neis, pnode.id)
				nei1Node.neis[nei2] = w1 + w2

				delete(nei2Node.neis, pnode.id)
				nei2Node.neis[nei1] = w1 + w2

				delete(g.nodes, pnode.id)

				dirty = true
				break
			}
		}
	}
}

func (g *Graph) PrintAsDot() string {
	ret := fmt.Sprintln("graph G {")

	for _, pnode := range g.nodes {
		ret += fmt.Sprintf("  %d\n", pnode.id)
		for nei, w := range pnode.neis {
			if nei > pnode.id {
				ret += fmt.Sprintf("    %d -- %d [label = %d]\n", pnode.id, nei, w)
			}
		}
	}

	ret += fmt.Sprintln("}")
	return ret
}
