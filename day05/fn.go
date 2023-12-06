package day05

import (
	"math"
	"slices"
)

type Fn struct {
	domain  []int
	offsets map[int]int // x -> offset
}

func NewFn(domain []int, offsets []int) *Fn {
	if len(domain) != len(offsets) {
		panic("400")
	}
	slices.Sort(domain)

	offsetsMap := make(map[int]int, len(domain))
	for i, d := range domain {
		offsetsMap[d] = offsets[i]
	}

	return &Fn{
		domain:  domain,
		offsets: offsetsMap,
	}
}

func (fn *Fn) AddPointFromRaw(dest, src, len int) {
	offset := dest - src
	oldOffset := fn.OffsetAt(src + len)
	fn.AddPoint(src, offset)
	// _, exists := fn.offsets[src+len]
	// if !exists {
	fn.AddPoint(src+len, oldOffset)
	// }
}

func (fn *Fn) AddPoint(x, offset int) {
	if x == math.MinInt || x == math.MaxInt {
		if offset != 0 {
			panic("offset must be 0 at -/+inf")
		}
		return // nop
	}

	_, exists := fn.offsets[x]
	if exists {
		// don't add to domain slice
		fn.offsets[x] = offset
		return
	}

	fn.domain = append(fn.domain, x)
	slices.Sort(fn.domain)
	fn.offsets[x] = offset
}

func (fn Fn) ValueAt(x int) (y int) {
	return fn.OffsetAt(x) + x
}

func (fn Fn) OffsetAt(x int) (offset int) {
	if len(fn.domain) == 0 {
		return 0
	}
	if x < fn.domain[0] {
		return 0
	}
	lower := fn.domainLowerEqual(x)
	return fn.offsets[lower[len(lower)-1]]
}

func (fn Fn) domainLowerEqual(value int) (ret []int) {
	for _, x := range fn.domain {
		if x > value {
			return
		}
		ret = append(ret, x)
	}
	return
}

func (fn Fn) domainGreaterEqual(value int) (ret []int) {
	for _, x := range fn.domain {
		if x > value {
			ret = append(ret, x)
		}
	}
	return
}

func (fn Fn) domainBetween(from, to int) (ret []int) {
	for _, x := range fn.domain {
		if x > to {
			return
		}

		if x >= from {
			ret = append(ret, x)
			continue
		}
	}
	return
}

// func (fn *Fn) Compress() {
// 	if len(fn.domain) <= 1 {
// 		return
// 	}

// 	for i := 0; i < len(fn.domain)-1; i++ {
// 		if fn.OffsetAt(fn.domain[i]) == fn.OffsetAt(fn.domain[i+1]) {
// 			// remove point at i+1
// 		}
// 	}
// }

func (fn Fn) ForDomain(from, to int) Fn {
	ret := *NewFn([]int{}, []int{})
	if len(fn.domain) == 0 {
		return ret
	}
	between := fn.domainBetween(from, to)
	if len(between) == 0 {
		ret.AddPoint(from, fn.OffsetAt(from))
		return ret
	}
	if from < between[0] {
		ret.AddPoint(from, fn.OffsetAt(from))
	}
	for _, b := range between {
		ret.AddPoint(b, fn.OffsetAt(b))
	}
	return ret
}

func Compose(f, g Fn) Fn {
	gf := *NewFn([]int{}, []int{})
	if len(f.domain) == 0 {
		return g
	}

	// for every interval of f domain
	fx := f.domain[0]
	fy := fx                                // offset 0
	lower := g.ForDomain(math.MinInt, fy-1) // open end interval
	if len(lower.domain) == 0 {
		// nop (offset is 0)
	} else {
		for _, gx := range lower.domain {
			fx = Inverse(gx, 0)
			gf.AddPoint(fx, g.OffsetAt(gx))
		}
	}

	if len(f.domain) > 1 {
		for i := 0; i < len(f.domain)-1; i++ {
			// compute the Fn of g
			fx_from := f.domain[i]
			fx_to := f.domain[i+1]
			offset := f.OffsetAt((fx_from + fx_to) / 2)
			fy_from := fx_from + offset
			fy_to := fx_to + offset

			between := g.ForDomain(fy_from, fy_to)
			for _, gx := range between.domain {
				fx = Inverse(gx, offset)
				gf.AddPoint(fx, offset+g.OffsetAt(gx))
			}
		}
	}

	fx = f.domain[len(f.domain)-1]
	if f.OffsetAt(fx) != 0 {
		panic("f should end with offset 0")
	}
	fy = fx
	upper := g.ForDomain(fx, math.MaxInt)
	for _, gx := range upper.domain {
		fx = Inverse(gx, 0)
		gf.AddPoint(fx, g.OffsetAt(gx))
	}

	//TODO at the end compress gf
	return gf
}

func Inverse(y int, offset int) (x int) {
	x = y - offset
	return
}
