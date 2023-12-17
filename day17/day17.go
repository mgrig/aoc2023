package day17

import (
	"fmt"
	"math"
)

const (
	N int = 0
	S int = 1
	E int = 2
	W int = 3
)

func Part1(lines []string) int {
	n := len(lines)
	g := NewGrid(n)

	for r, line := range lines {
		for c, run := range line {
			g.grid[r][c] = int(run) - int('0')
		}
	}
	fmt.Println(g)

	states := make(map[State]int, 0)
	startState := NewState(0, 0, E, 3)
	toVisit := []State{startState}
	visited := make(map[State]bool)
	states[startState] = g.grid[0][0]

	for len(toVisit) > 0 {
		dijkstra(g, &states, &toVisit, &visited)
	}

	minDist := math.MaxInt
	for state, dist := range states {
		if state.r == n-1 && state.c == n-1 && dist < minDist {
			minDist = dist
		}
	}

	return minDist - g.grid[0][0]
}

func dijkstra(g *Grid, states *map[State]int, toVisit *[]State, visited *map[State]bool) {
	state := getNextToVisit(states, toVisit)

	// find possible next steps (keys)
	nextStates := make([]State, 0)
	if state.maxStepsInDir > 0 {
		addToNextStates(g, &nextStates, state.NextInDir())
	}
	addToNextStates(g, &nextStates, state.NextRight())
	addToNextStates(g, &nextStates, state.NextLeft())

	n := len(g.grid)
	if state.r == n-1 && state.c == n-1 {
		return
	}

	for _, nextState := range nextStates {
		possibleValue(states, nextState, (*states)[state]+g.grid[nextState.r][nextState.c])
		addToVisit(toVisit, visited, nextState)
	}
	(*visited)[state] = true
}

func getNextToVisit(states *map[State]int, toVisit *[]State) (nextState State) {
	var bestIndex int
	min := math.MaxInt
	for i, st := range *toVisit {
		if (*states)[st] < min {
			min = (*states)[st]
			bestIndex = i
		}
	}
	nextState = (*toVisit)[bestIndex]
	*toVisit = append((*toVisit)[:bestIndex], (*toVisit)[bestIndex+1:]...)
	return nextState
}

func addToVisit(toVisit *[]State, visited *map[State]bool, state State) {
	_, exists := (*visited)[state]
	if exists {
		return
	}
	for _, oldState := range *toVisit {
		if oldState == state {
			return
		}
	}
	*toVisit = append(*toVisit, state)
}

func possibleValue(states *map[State]int, state State, value int) {
	oldValue, exists := (*states)[state]
	if !exists || value < oldValue {
		(*states)[state] = value
	}
}

func addToNextStates(g *Grid, nextSteps *[]State, key State) {
	if key.maxStepsInDir >= 0 && g.Inside(key.r, key.c) {
		*nextSteps = append(*nextSteps, key)
	}
}
