package day05

type DomainInterval struct {
	from, to int
}

func NewDomainInterval(from, to int) *DomainInterval {
	return &DomainInterval{
		from: from,
		to:   to,
	}
}

func (di *DomainInterval) Contains(value int) bool {
	return value >= di.from && value <= di.to
}

// ***

type DomainIntervals []DomainInterval

func (dis *DomainIntervals) Contains(value int) bool {
	for _, di := range *dis {
		if di.Contains(value) {
			return true
		}
	}
	return false
}
