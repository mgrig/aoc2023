package day05

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
}

func (m *Mapping) Get(source int) (dest int) {
	for _, rangeMap := range m.rangeMaps {
		if rangeMap.Contains(source) {
			return rangeMap.Get(source)
		}
	}
	return source
}
