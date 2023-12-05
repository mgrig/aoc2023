package day05

import (
	"aoc2023/common"
	"fmt"
)

type RangeMap struct {
	sourceStart, destStart, len int
}

func NewRangeMap(source, dest, len int) *RangeMap {
	return &RangeMap{
		sourceStart: source,
		destStart:   dest,
		len:         len,
	}
}

func (rm *RangeMap) Contains(value int) bool {
	return value >= rm.sourceStart && value < rm.sourceStart+rm.len
}

func (rm *RangeMap) SourceEnd() int {
	return rm.sourceStart + rm.len - 1
}

func (rm *RangeMap) Get(value int) (dest int) {
	if !rm.Contains(value) {
		panic("Get with value outside source range!")
	}
	return (value - rm.sourceStart) + rm.destStart
}

// ****

type Mapping struct {
	rangeMaps []RangeMap
}

func NewMapping() *Mapping {
	return &Mapping{
		rangeMaps: make([]RangeMap, 0),
	}
}

func (m *Mapping) AddRange(rangeMap RangeMap) {
	m.rangeMaps = append(m.rangeMaps, rangeMap)
	prev := len(m.rangeMaps)
	m.NormalizeRanges()
	next := len(m.rangeMaps)
	if next != prev {
		fmt.Println("normalize ", prev, " > ", next)
	}
}

// Merge ranges if possible.
// Should result in a canonical form (impossible to reduce further).
// Changes are performed in-place.
func (m *Mapping) NormalizeRanges() {
	if len(m.rangeMaps) <= 1 {
		return
	}

	dirty := true
	for dirty {
		dirty = false

		for i := 0; i < len(m.rangeMaps) && !dirty; i++ {
			for j := i + 1; j < len(m.rangeMaps) && !dirty; j++ {
				if rangesOverlap(&m.rangeMaps[i], &m.rangeMaps[j]) {
					dirty = true
					mergedRange := mergeRanges(&m.rangeMaps[i], &m.rangeMaps[j])
					m.rangeMaps[i] = *mergedRange

					// delete element j
					m.rangeMaps[j] = m.rangeMaps[len(m.rangeMaps)-1] // replace j with last element
					m.rangeMaps = m.rangeMaps[:len(m.rangeMaps)-1]   // keep only len-1 elements
				}
			}
		}
	}
}

func (m *Mapping) Get(source int) (dest int) {
	for _, rangeMap := range m.rangeMaps {
		if rangeMap.Contains(source) {
			return rangeMap.Get(source)
		}
	}
	return source
}

func mergeRanges(rm1, rm2 *RangeMap) *RangeMap {
	if !rangesOverlap(rm1, rm2) {
		panic("ranges must overlap")
	}
	from := common.IntMin(rm1.sourceStart, rm2.sourceStart)
	to := common.IntMax(rm1.SourceEnd(), rm2.SourceEnd())
	len := to - from + 1
	return NewRangeMap(from, from, len)
}

// work only on source, i.e. ignore dest
func rangesOverlap(rm1, rm2 *RangeMap) bool {
	return rm1.sourceStart <= rm2.SourceEnd() && rm1.SourceEnd() >= rm2.sourceStart
}
