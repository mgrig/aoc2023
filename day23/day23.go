package day23

import (
	"aoc2023/common"
	"fmt"
)

func Part1(lines []string) int {
	n := len(lines)
	start := NewCoord(0, findFirstEmpty(lines[0]))
	// end := NewCoord(n-1, findFirstEmpty(lines[n-1]))
	// fmt.Println(start, end)

	propagated := map[Coord]int{}
	toPropagate := map[Coord]int{start: 0}
	nodes := NewNodes()
	nodes.Set(*NewNode(0))

	for len(toPropagate) > 0 {
		var pos Coord
		var nodeId int
		for pos, nodeId = range toPropagate {
			break
		}
		delete(toPropagate, pos)
		if _, exists := propagated[pos]; exists {
			continue
		}
		propagated[pos] = nodeId

		size := 1
		// follow path
		visited := map[Coord]bool{}
		visited[pos] = true
		for {
			neighbors := []Coord{NewCoord(pos.r+1, pos.c), NewCoord(pos.r, pos.c+1), NewCoord(pos.r-1, pos.c), NewCoord(pos.r, pos.c-1)}
			neiFound := false
			for _, nei := range neighbors {
				_, alreadyVisited := visited[nei]
				if inside(nei, n) && !alreadyVisited && lines[nei.r][nei.c] == '.' {
					visited[nei] = true
					pos = nei
					size++
					neiFound = true
					break
				}
			}

			if !neiFound { // end of path
				node := nodes.GetOrCreate(nodeId)
				node.size = size

				down := neighbors[0]
				if inside(down, n) && !visited[down] && lines[down.r][down.c] == 'v' {
					nextPos := NewCoord(down.r+1, down.c)
					nextId, exists := getPropagatedOrPlanned(&propagated, &toPropagate, nextPos)
					if !exists {
						nextId = len(*nodes)
					}
					nodes.GetOrCreate(nextId)
					node.AddNext(nextId)
					toPropagate[nextPos] = nextId
				}

				right := neighbors[1]
				if inside(right, n) && !visited[right] && lines[right.r][right.c] == '>' {
					nextPos := NewCoord(right.r, right.c+1)
					nextId, exists := getPropagatedOrPlanned(&propagated, &toPropagate, nextPos)
					if !exists {
						nextId = len(*nodes)
					}
					nodes.GetOrCreate(nextId)
					node.AddNext(nextId)
					toPropagate[nextPos] = nextId
				}

				up := neighbors[2]
				if inside(up, n) && !visited[up] && lines[up.r][up.c] == '^' {
					nextPos := NewCoord(up.r-1, up.c)
					nextId, exists := getPropagatedOrPlanned(&propagated, &toPropagate, nextPos)
					if !exists {
						nextId = len(*nodes)
					}
					nodes.GetOrCreate(nextId)
					node.AddNext(nextId)
					toPropagate[nextPos] = nextId
				}

				left := neighbors[3]
				if inside(left, n) && !visited[left] && lines[left.r][left.c] == '<' {
					nextPos := NewCoord(left.r, left.c-1)
					nextId, exists := getPropagatedOrPlanned(&propagated, &toPropagate, nextPos)
					if !exists {
						nextId = len(*nodes)
					}
					nodes.GetOrCreate(nextId)
					node.AddNext(nextId)
					toPropagate[nextPos] = nextId
				}

				nodes.Set(node)
				break
			}
		}
	}

	// for _, node := range *nodes {
	// 	fmt.Printf("  %d[label=\"%d(%d)\"]\n", node.id, node.id, node.size)
	// 	for _, next := range node.next {
	// 		fmt.Printf("    %d -> %d\n", node.id, next)
	// 	}
	// }

	// search longest path
	startNodeId := 0
	endNodeId := nodes.GetEndNode().id
	// fmt.Println(startNodeId, endNodeId)

	dist := make(map[int]int, 0)
	for nodeId := range *nodes {
		dist[nodeId] = 0
	}
	dist[startNodeId] = (*nodes)[startNodeId].size

	toVisit := []int{startNodeId}
	for len(toVisit) > 0 {
		nodeId := toVisit[0]
		toVisit = toVisit[1:]

		node := (*nodes)[nodeId]
		for _, nextId := range node.next {
			next := (*nodes)[nextId]
			dist[nextId] = common.IntMax(dist[nextId], dist[nodeId]+1+next.size)
		}
		toVisit = append(toVisit, node.next...)
	}

	// result for part 1: dist[endNodeId] - 1

	g := NewGraph()
	// transform graph to a new, undirected one, for part 2
	for _, node := range *nodes {
		// fmt.Printf("  %d\n", node.id)
		for _, nextId := range node.next {
			next := (*nodes)[nextId]
			// fmt.Printf("    %d -- %d [label=%d]\n", node.id, nextId, 1+next.size)
			g.AddEdge(node.id, nextId, 1+next.size)
		}
	}
	g.AddEdge(0, 1, g.GetWeight(0, 1)+dist[startNodeId]-1)

	// fmt.Println(g.PrintAsDot())
	g.Simplify()
	// fmt.Println("simplified:")
	fmt.Println(g.PrintAsDot())

	// search all paths from start to end, without revisit
	visited := make(map[int]bool, 0)
	maxLen := 0
	rec(g, visited, startNodeId, 0, endNodeId, &maxLen)

	return maxLen
}

func rec(g *Graph, visited map[int]bool, currentNodeId, currentLen, endNodeId int, maxLen *int) {
	if currentNodeId == endNodeId {
		*maxLen = common.IntMax(*maxLen, currentLen)
	}

	currentNode := g.GetOrCreateUNode(currentNodeId)

	visitedClone := make(map[int]bool, len(visited))
	for k, v := range visited {
		visitedClone[k] = v
	}
	visitedClone[currentNodeId] = true

	for nei, w := range currentNode.neis {
		if _, exists := visited[nei]; !exists {
			rec(g, visitedClone, nei, currentLen+w, endNodeId, maxLen)
		}
	}
}

func findFirstEmpty(line string) (index int) {
	for c, val := range line {
		if val == '.' {
			return c
		}
	}
	panic("empty not found")
}

func inside(pos Coord, n int) bool {
	return pos.r >= 0 && pos.c >= 0 && pos.r < n && pos.c < n
}

func getPropagatedOrPlanned(propagated *map[Coord]int, toPropagate *map[Coord]int, pos Coord) (nodeId int, found bool) {
	nodeId, exists := (*propagated)[pos]
	if exists {
		return nodeId, true
	}

	nodeId, exists = (*toPropagate)[pos]
	if exists {
		return nodeId, true
	}

	return -1, false
}
