package day19

type Range4 struct {
	ranges map[string]Range1 // [xmas] -> Range1
}

func NewRange4() Range4 {
	return Range4{
		ranges: map[string]Range1{
			X: NewRange1(1, 4000),
			M: NewRange1(1, 4000),
			A: NewRange1(1, 4000),
			S: NewRange1(1, 4000),
		},
	}
}

func (r4 Range4) Prod() int {
	prod := 1
	for _, v := range r4.ranges {
		prod *= v.Count()
	}
	return prod
}

func (r4 Range4) Inside(category string, splitValue int) bool {
	return r4.ranges[category].Inside(splitValue)
}

func (r4 Range4) SplitSmallerThan(category string, splitValue int) (left4, right4 Range4) {
	left4 = r4.Clone()
	right4 = r4.Clone()

	left, right := r4.ranges[category].SplitSmallerThan(splitValue)
	left4.ranges[category] = left
	right4.ranges[category] = right
	return
}

func (r4 Range4) Clone() Range4 {
	ret := NewRange4()
	for k, v := range r4.ranges {
		ret.ranges[k] = v
	}
	return ret
}

// ****

type Range1 struct {
	min, max int
}

func NewRange1(min, max int) Range1 {
	if min > max {
		panic("wrong min/max")
	}
	return Range1{
		min: min,
		max: max,
	}
}

func (r1 Range1) Count() int {
	return r1.max - r1.min + 1
}

func (r1 Range1) Inside(value int) bool {
	return value >= r1.min && value <= r1.max
}

// [min, max] -> [min, split-1] + [split, max]
func (r1 Range1) SplitSmallerThan(value int) (left, right Range1) {
	if !r1.Inside(value) {
		panic("wrong split")
	}
	left = NewRange1(r1.min, value-1)
	right = NewRange1(value, r1.max)
	return
}
