package day05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFn_domainLowerEqual(t *testing.T) {
	fn := NewFn([]int{2, 5, 7}, []int{3, -4, 0})
	lower := fn.domainLowerEqual(6)
	assert.Equal(t, lower, []int{2, 5})
}

func TestFn_between(t *testing.T) {
	fn := NewFn([]int{2, 5, 7}, []int{3, -4, 0})
	lower := fn.domainBetween(2, 6)
	assert.Equal(t, lower, []int{2, 5})
}

func Test_Compose(t *testing.T) {
	f := NewFn([]int{2, 4}, []int{3, 0})
	g := NewFn([]int{1, 3}, []int{-4, 0})
	gf := Compose(*f, *g)
	fmt.Println(gf)
}

func Test_ComposeTHL(t *testing.T) {
	// f := NewFn([]int{0, 69, 70}, []int{1, -69, 0})  // t 2 h
	// g := NewFn([]int{56, 93, 97}, []int{4, -37, 0}) // h 2 l

	f := NewFn([]int{}, []int{})
	f.AddPointFromRaw(0, 69, 1)
	f.AddPointFromRaw(1, 0, 69)

	g := NewFn([]int{}, []int{})
	g.AddPointFromRaw(60, 56, 37)
	g.AddPointFromRaw(56, 93, 4)
	fmt.Println(*f, *g)

	gf := Compose(*f, *g)
	fmt.Println(gf)
}
